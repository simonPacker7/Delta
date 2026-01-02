package redisclient

import (
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/simonPacker7/Delta/backend/shared/entities"
)

const gameKeyPrefix = "game:"
const matchmakingQueue = "game:open"
const privateGameCodePrefix = "game:code:"

// CreateGame creates a new game in Redis with status "waiting"
func (r *RedisClient) CreateGame(game entities.Game) error {
	key := gameKeyPrefix + game.ID

	err := r.client.HSet(ctx, key,
		"id", game.ID,
		"type", string(game.Type),
		"status", string(game.Status),
		"join_code", game.JoinCode,
		"player1_id", game.Player1ID,
		"player1_name", game.Player1Name,
		"player2_id", game.Player2ID,
		"player2_name", game.Player2Name,
		"current_word", game.CurrentWord,
		"current_turn_id", game.CurrentTurnID,
		"connected_count", game.ConnectedCount,
		"created_at", game.CreatedAt,
	).Err()

	if err != nil {
		return err
	}

	// Set TTL of 24 hours
	return r.client.Expire(ctx, key, 24*time.Hour).Err()
}

// GetGame retrieves a game from Redis by ID
func (r *RedisClient) GetGame(gameID string) (entities.Game, error) {
	key := gameKeyPrefix + gameID

	var game entities.Game
	err := r.client.HGetAll(ctx, key).Scan(&game)
	if err != nil {
		return entities.Game{}, err
	}

	return game, nil
}

// UpdateGame updates specific fields of a game
func (r *RedisClient) UpdateGame(gameID string, fields map[string]interface{}) error {
	key := gameKeyPrefix + gameID
	return r.client.HSet(ctx, key, fields).Err()
}

// DoesGameExist checks if a game exists in Redis
func (r *RedisClient) DoesGameExist(gameID string) (bool, error) {
	key := gameKeyPrefix + gameID
	return r.DoesRecordExist(key)
}

// PushToMatchmakingQueue adds a game ID to the matchmaking queue
func (r *RedisClient) PushToMatchmakingQueue(gameID string) error {
	return r.client.LPush(ctx, matchmakingQueue, gameID).Err()
}

// PopFromMatchmakingQueue removes and returns a game ID from the matchmaking queue
func (r *RedisClient) PopFromMatchmakingQueue() (string, error) {
	result, err := r.client.RPop(ctx, matchmakingQueue).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

// PublishGameEvent publishes an event to the game's pub/sub channel
func (r *RedisClient) PublishGameEvent(gameID string, event string) error {
	channel := gameKeyPrefix + gameID
	return r.client.Publish(ctx, channel, event).Err()
}

// SetPrivateGameCode maps a join code to a game ID with TTL
func (r *RedisClient) SetPrivateGameCode(joinCode string, gameID string) error {
	key := privateGameCodePrefix + joinCode
	err := r.client.Set(ctx, key, gameID, 24*time.Hour).Err()
	return err
}

// GetGameIDByCode retrieves a game ID by its join code
func (r *RedisClient) GetGameIDByCode(joinCode string) (string, error) {
	key := privateGameCodePrefix + joinCode
	return r.client.Get(ctx, key).Result()
}

// DeletePrivateGameCode removes the join code mapping after use
func (r *RedisClient) DeletePrivateGameCode(joinCode string) error {
	key := privateGameCodePrefix + joinCode
	return r.client.Del(ctx, key).Err()
}

// AtomicJoinPrivateGame atomically validates and joins a private game
// Returns: gameID, player1ID, error
// Uses Lua script to ensure only one player can join
var joinPrivateGameScript = `
local codeKey = KEYS[1]
local gameId = redis.call('GET', codeKey)
if not gameId then
    return {'', '', 'invalid_code'}
end

local gameKey = 'game:' .. gameId
local wordsKey = gameKey .. ':words'
local status = redis.call('HGET', gameKey, 'status')
if status ~= 'waiting' then
    return {'', '', 'game_not_available'}
end

local player1Id = redis.call('HGET', gameKey, 'player1_id')
local player2Id = ARGV[1]
local player2Name = ARGV[2]
local startWord = ARGV[3]

if player1Id == player2Id then
    return {'', '', 'cannot_join_own_game'}
end

-- Atomically update game and delete code
redis.call('HSET', gameKey, 
    'player2_id', player2Id,
    'player2_name', player2Name,
    'status', 'ready',
    'current_word', startWord,
    'current_turn_id', player1Id
)

-- Initialize played words set with starting word
redis.call('SADD', wordsKey, startWord)
redis.call('EXPIRE', wordsKey, 86400)

redis.call('DEL', codeKey)

return {gameId, player1Id, ''}
`

func (r *RedisClient) AtomicJoinPrivateGame(joinCode string, player2ID string, player2Name string, startWord string) (string, string, error) {
	codeKey := privateGameCodePrefix + joinCode

	result, err := r.client.Eval(ctx, joinPrivateGameScript, []string{codeKey}, player2ID, player2Name, startWord).Result()
	if err != nil {
		return "", "", err
	}

	// Parse result array
	arr, ok := result.([]interface{})
	if !ok || len(arr) != 3 {
		return "", "", err
	}

	gameID, _ := arr[0].(string)
	player1ID, _ := arr[1].(string)
	errMsg, _ := arr[2].(string)

	if errMsg != "" {
		return "", "", &AtomicOperationError{Message: errMsg}
	}

	return gameID, player1ID, nil
}

// AtomicPopAndJoinGame atomically pops from queue and joins if valid
// Returns: gameID, player1ID, matched (bool), error
var popAndJoinGameScript = `
local queueKey = KEYS[1]
local playerId = ARGV[1]
local playerName = ARGV[2]
local startWord = ARGV[3]
local maxAttempts = 10

for i = 1, maxAttempts do
    local gameId = redis.call('RPOP', queueKey)
    if not gameId then
        return {'', '', false}
    end

    local gameKey = 'game:' .. gameId
    local wordsKey = gameKey .. ':words'
    local status = redis.call('HGET', gameKey, 'status')
    local player1Id = redis.call('HGET', gameKey, 'player1_id')

    -- Skip invalid games
    if status == 'waiting' and player1Id and player1Id ~= '' then
        -- Don't match with self - put back and return not found
        if player1Id == playerId then
            redis.call('LPUSH', queueKey, gameId)
            return {'', '', false}
        end

        -- Valid game - join it
        redis.call('HSET', gameKey,
            'player2_id', playerId,
            'player2_name', playerName,
            'status', 'ready',
            'current_word', startWord,
            'current_turn_id', player1Id
        )
        
        -- Initialize played words set with starting word
        redis.call('SADD', wordsKey, startWord)
        redis.call('EXPIRE', wordsKey, 86400)
        
        return {gameId, player1Id, true}
    end
end

return {'', '', false}
`

func (r *RedisClient) AtomicPopAndJoinGame(playerID string, playerName string, startWord string) (string, string, bool, error) {
	result, err := r.client.Eval(ctx, popAndJoinGameScript, []string{matchmakingQueue}, playerID, playerName, startWord).Result()
	if err != nil {
		return "", "", false, err
	}

	arr, ok := result.([]interface{})
	if !ok || len(arr) != 3 {
		return "", "", false, nil
	}

	gameID, _ := arr[0].(string)
	player1ID, _ := arr[1].(string)
	matched := false
	if m, ok := arr[2].(int64); ok {
		matched = m == 1
	}

	return gameID, player1ID, matched, nil
}

// AtomicOperationError represents an error from an atomic operation
type AtomicOperationError struct {
	Message string
}

func (e *AtomicOperationError) Error() string {
	return e.Message
}

// ========== Game Expiration Operations ==========

const gameExpireSet = "game:expire"
const turnTimeoutSeconds = 100

// AddGameToExpireQueue adds a game to the expiration sorted set
// Score is current timestamp + timeout seconds
func (r *RedisClient) AddGameToExpireQueue(gameID string) error {
	expireAt := time.Now().Unix() + turnTimeoutSeconds
	return r.client.ZAdd(ctx, gameExpireSet, redis.Z{
		Score:  float64(expireAt),
		Member: gameID,
	}).Err()
}

// AtomicUpdateGameExpiration atomically removes and re-adds a game to the expiration queue
// Used when a player makes a valid move to reset the turn timer
var updateGameExpirationScript = `
local expireSet = KEYS[1]
local gameId = ARGV[1]
local newExpireTime = ARGV[2]

redis.call('ZREM', expireSet, gameId)
redis.call('ZADD', expireSet, newExpireTime, gameId)
return 1
`

func (r *RedisClient) AtomicUpdateGameExpiration(gameID string) error {
	expireAt := time.Now().Unix() + turnTimeoutSeconds
	_, err := r.client.Eval(ctx, updateGameExpirationScript, []string{gameExpireSet}, gameID, expireAt).Result()
	return err
}

// RemoveGameFromExpireQueue removes a game from the expiration queue (when game ends normally)
func (r *RedisClient) RemoveGameFromExpireQueue(gameID string) error {
	return r.client.ZRem(ctx, gameExpireSet, gameID).Err()
}

// AtomicClaimAndEndExpiredGames atomically claims expired games and ends them
// Returns a list of ended games with their winners
// This prevents race conditions where a player moves between claim and end
var claimAndEndExpiredGamesScript = `
local expireSet = KEYS[1]
local gamePrefix = 'game:'
local now = ARGV[1]
local limit = ARGV[2]

local expired = redis.call('ZRANGEBYSCORE', expireSet, 0, now, 'LIMIT', 0, limit)
if #expired == 0 then
    return {}
end

local results = {}

for i, gameId in ipairs(expired) do
    local gameKey = gamePrefix .. gameId
    local status = redis.call('HGET', gameKey, 'status')
    
    -- Only end games that are still active (not already ended)
    if status == 'active' then
        local currentTurnId = redis.call('HGET', gameKey, 'current_turn_id')
        local player1Id = redis.call('HGET', gameKey, 'player1_id')
        local player2Id = redis.call('HGET', gameKey, 'player2_id')
        
        -- Winner is the player who was NOT the current turn
        local winnerId = player1Id
        if currentTurnId == player1Id then
            winnerId = player2Id
        end
        
        -- Atomically end the game
        redis.call('HSET', gameKey, 
            'status', 'completed',
            'winner_id', winnerId,
            'win_reason', 'timeout'
        )
        
        -- Remove from expire set
        redis.call('ZREM', expireSet, gameId)
        
        -- Publish JSON event for clients
        local jsonMsg = '{"type":"game_ended","gameId":"' .. gameId .. '","payload":{"winnerId":"' .. winnerId .. '","reason":"timeout"}}'
        redis.call('PUBLISH', gameKey, jsonMsg)
        
        -- Add to results
        table.insert(results, {gameId, winnerId})
    else
        -- Game already ended or not active, just remove from expire set
        redis.call('ZREM', expireSet, gameId)
    end
end

return results
`

// ExpiredGameResult represents a game that was ended due to timeout
type ExpiredGameResult struct {
	GameID   string
	WinnerID string
}

func (r *RedisClient) AtomicClaimAndEndExpiredGames(limit int) ([]ExpiredGameResult, error) {
	now := time.Now().Unix()
	result, err := r.client.Eval(ctx, claimAndEndExpiredGamesScript, []string{gameExpireSet}, now, limit).Result()
	if err != nil {
		return nil, err
	}

	// Parse result array of [gameId, winnerId] pairs
	arr, ok := result.([]interface{})
	if !ok {
		return []ExpiredGameResult{}, nil
	}

	results := make([]ExpiredGameResult, 0, len(arr))
	for _, item := range arr {
		if pair, ok := item.([]interface{}); ok && len(pair) == 2 {
			gameID, _ := pair[0].(string)
			winnerID, _ := pair[1].(string)
			results = append(results, ExpiredGameResult{
				GameID:   gameID,
				WinnerID: winnerID,
			})
		}
	}

	return results, nil
}

// Deprecated: Use AtomicClaimAndEndExpiredGames instead
// AtomicClaimExpiredGames atomically gets and removes expired games from the queue
var claimExpiredGamesScript = `
local expireSet = KEYS[1]
local now = ARGV[1]
local limit = ARGV[2]

local expired = redis.call('ZRANGEBYSCORE', expireSet, 0, now, 'LIMIT', 0, limit)
if #expired > 0 then
    redis.call('ZREM', expireSet, unpack(expired))
end
return expired
`

func (r *RedisClient) AtomicClaimExpiredGames(limit int) ([]string, error) {
	now := time.Now().Unix()
	result, err := r.client.Eval(ctx, claimExpiredGamesScript, []string{gameExpireSet}, now, limit).Result()
	if err != nil {
		return nil, err
	}

	arr, ok := result.([]interface{})
	if !ok {
		return []string{}, nil
	}

	gameIDs := make([]string, 0, len(arr))
	for _, item := range arr {
		if id, ok := item.(string); ok {
			gameIDs = append(gameIDs, id)
		}
	}

	return gameIDs, nil
}

// EndGameByTimeout updates a game's status to completed due to timeout
// Deprecated: Use AtomicClaimAndEndExpiredGames instead for atomic operation
func (r *RedisClient) EndGameByTimeout(gameID string, winnerID string) error {
	return r.client.HSet(ctx, gameKeyPrefix+gameID,
		"status", string(entities.GameStatusEnded),
		"winner_id", winnerID,
		"win_reason", "timeout",
	).Err()
}

// AtomicJoinGameSession atomically increments connected_count and starts game if both players connected
// Returns: newConnectedCount, gameStarted, error
var joinGameSessionScript = `
local gameKey = KEYS[1]
local expireSet = KEYS[2]
local turnTimeout = ARGV[1]

local status = redis.call('HGET', gameKey, 'status')
if not status then
    return {-1, false, 'game_not_found'}
end

-- Allow joining games in waiting (player1 before match), ready (both matched), or active (reconnect)
if status ~= 'waiting' and status ~= 'ready' and status ~= 'active' then
    return {-1, false, 'game_not_joinable'}
end

local newCount = redis.call('HINCRBY', gameKey, 'connected_count', 1)

-- Only start the game when:
-- 1. Both players are connected (count == 2)
-- 2. Status is 'ready' (both players have been matched via matchmaking)
if newCount == 2 and status == 'ready' then
    redis.call('HSET', gameKey, 'status', 'active')
    
    -- Add to expiration queue
    local timeResult = redis.call('TIME')
    local expireAt = tonumber(timeResult[1]) + tonumber(turnTimeout)
    local gameId = string.gsub(gameKey, 'game:', '')
    redis.call('ZADD', expireSet, expireAt, gameId)
    
    return {newCount, true, ''}
end

return {newCount, false, ''}
`

func (r *RedisClient) AtomicJoinGameSession(gameID string) (int, bool, error) {
	gameKey := gameKeyPrefix + gameID
	result, err := r.client.Eval(ctx, joinGameSessionScript, []string{gameKey, gameExpireSet}, turnTimeoutSeconds).Result()
	if err != nil {
		return 0, false, err
	}

	arr, ok := result.([]interface{})
	if !ok || len(arr) != 3 {
		return 0, false, nil
	}

	count, _ := arr[0].(int64)
	started := false
	if s, ok := arr[1].(int64); ok {
		started = s == 1
	}
	errMsg, _ := arr[2].(string)

	if errMsg != "" {
		return 0, false, &AtomicOperationError{Message: errMsg}
	}

	return int(count), started, nil
}

// AtomicLeaveGameSession decrements connected_count when a player disconnects
func (r *RedisClient) AtomicLeaveGameSession(gameID string) error {
	gameKey := gameKeyPrefix + gameID
	return r.client.HIncrBy(ctx, gameKey, "connected_count", -1).Err()
}

// AtomicSubmitWord atomically validates and applies a word submission
// Returns: success, newWord, nextTurnPlayerID, error
// Also tracks played words in a set to prevent duplicates
var submitWordScript = `
local gameKey = KEYS[1]
local expireSet = KEYS[2]
local wordsKey = KEYS[3]
local playerId = ARGV[1]
local newWord = ARGV[2]
local turnTimeout = ARGV[3]

local status = redis.call('HGET', gameKey, 'status')
if status ~= 'active' then
    return {false, '', '', 'game_not_active'}
end

local currentTurnId = redis.call('HGET', gameKey, 'current_turn_id')
if currentTurnId ~= playerId then
    return {false, '', '', 'not_your_turn'}
end

-- Check if word has already been played
if redis.call('SISMEMBER', wordsKey, newWord) == 1 then
    return {false, '', '', 'word_already_played'}
end

-- Determine next turn player
local player1Id = redis.call('HGET', gameKey, 'player1_id')
local player2Id = redis.call('HGET', gameKey, 'player2_id')
local nextTurnId = player1Id
if playerId == player1Id then
    nextTurnId = player2Id
end

-- Add word to played words set
redis.call('SADD', wordsKey, newWord)

-- Update game state
redis.call('HSET', gameKey, 
    'current_word', newWord,
    'current_turn_id', nextTurnId
)

-- Reset turn timer
local timeResult = redis.call('TIME')
local expireAt = tonumber(timeResult[1]) + tonumber(turnTimeout)
local gameId = string.gsub(gameKey, 'game:', '')
redis.call('ZREM', expireSet, gameId)
redis.call('ZADD', expireSet, expireAt, gameId)

return {true, newWord, nextTurnId, ''}
`

func (r *RedisClient) AtomicSubmitWord(gameID string, playerID string, newWord string) (bool, string, string, error) {
	gameKey := gameKeyPrefix + gameID
	wordsKey := gameKeyPrefix + gameID + ":words"
	result, err := r.client.Eval(ctx, submitWordScript, []string{gameKey, gameExpireSet, wordsKey}, playerID, newWord, turnTimeoutSeconds).Result()
	if err != nil {
		return false, "", "", err
	}

	arr, ok := result.([]interface{})
	if !ok || len(arr) != 4 {
		return false, "", "", nil
	}

	success := false
	if s, ok := arr[0].(int64); ok {
		success = s == 1
	}
	word, _ := arr[1].(string)
	nextTurnID, _ := arr[2].(string)
	errMsg, _ := arr[3].(string)

	if errMsg != "" {
		return false, "", "", &AtomicOperationError{Message: errMsg}
	}

	return success, word, nextTurnID, nil
}

// GetPlayedWords retrieves all words that have been played in a game
func (r *RedisClient) GetPlayedWords(gameID string) ([]string, error) {
	wordsKey := gameKeyPrefix + gameID + ":words"
	return r.client.SMembers(ctx, wordsKey).Result()
}

// IsWordPlayed checks if a specific word has already been played in a game
func (r *RedisClient) IsWordPlayed(gameID string, word string) (bool, error) {
	wordsKey := gameKeyPrefix + gameID + ":words"
	return r.client.SIsMember(ctx, wordsKey, word).Result()
}
