import axios from "axios";
import { defineStore } from "pinia";
import { ref } from "vue";
import { useSocketStore } from "./socket";

export type GamePlayer = {
    Name: string,
    ID: string,
}

export type GameTurn = {
    Player: GamePlayer,
    Word: string,
    Timestamp: number,
}

// TODO
// 1. Get matchmaking working
// 2. Get basic game state information on join 
// 3. Handle sending words
// 4. Handle receiving game state updates from server

export const useGameStore = defineStore("game", () => {
    const gameId = ref<string | null>(null);
    const gameType = ref<string | null>(null);
    const gameStatus = ref<string | null>(null);
    const currentWord = ref<string | null>(null);
    const currentTurnId = ref<string | null>(null);
    const player1 = ref<GamePlayer | null>(null);
    const player2 = ref<GamePlayer | null>(null);
    const gameTurns = ref<GameTurn[]>([]);
    const winnerId = ref<string | null>(null);
    const winReason = ref<string>('');
    const joinCode = ref<string>('');
    const error = ref<string | null>(null);

    const $reset = () => {
        gameId.value = null;
        gameType.value = null;
        gameStatus.value = null;
        currentWord.value = null;
        currentTurnId.value = null;
        player1.value = null;
        player2.value = null;
        gameTurns.value = [];
        winnerId.value = null;
        winReason.value = '';
        joinCode.value = '';
        error.value = null;
    }

    const FindGame = async () => {
        try {
            const res = await axios.get('/api/game/find');
            if (!res || res.status != 200 || !res.data || !res.data.gameId) {
                $reset();
                return;
            }
            gameId.value = res.data.gameId;
            gameType.value = 'online';
            gameStatus.value = res.data.status; // "waiting" or "matched"
            SubscribeToGame(res.data.gameId);
        }
        catch (err) {
            $reset();
            console.log(err);
        }
    }

    const CreatePrivateGame = async () => {
        try {
            const res = await axios.post('/api/game/private/create');
            if (!res || res.status != 200 || !res.data || !res.data.gameId) {
                $reset();
                return;
            }
            gameId.value = res.data.gameId;
            joinCode.value = res.data.joinCode;
            gameType.value = 'private';
            gameStatus.value = 'waiting';
            SubscribeToGame(res.data.gameId);
        }
        catch (err) {
            $reset();
            console.log(err);
        }
    }

    const JoinPrivateGame = async (code: string) => {
        try {
            const res = await axios.post('/api/game/private/join', { joinCode: code });
            if (!res || res.status != 200 || !res.data || !res.data.gameId) {
                $reset();
                return;
            }
            gameId.value = res.data.gameId;
            gameType.value = 'private';
            gameStatus.value = 'ready';
            SubscribeToGame(res.data.gameId);
        }
        catch (err) {
            $reset();
            console.log(err);
        }
    }

    const SubscribeToGame = async (id: string) => {
        useSocketStore().sendSocketMessage(JSON.stringify({
            action: "join_game",
            gameId: id
        }));
    }

    const SendWord = async (word: string) => {
        useSocketStore().sendSocketMessage(JSON.stringify({
            action: "submit_word",
            gameId: gameId.value,
            word: word
        }));
    }

    const HandleGameUpdate = (data: string) => {
        console.log("Handling game update:", data);
        
        let update;
        try {
            update = JSON.parse(data);
        } catch (e) {
            console.error("Failed to parse game update:", e);
            return;
        }

        const type = update.type as string;
        const payload = update.payload;

        switch (type) {
            case "joined_game":
                // Confirmation that we joined the game
                console.log("Joined game:", update.gameId);
                break;

            case "game_started":
                // Both players connected, game is starting
                gameStatus.value = "active";
                currentWord.value = payload.currentWord;
                currentTurnId.value = payload.currentTurnId;
                player1.value = { ID: payload.player1Id, Name: payload.player1Name };
                player2.value = { ID: payload.player2Id, Name: payload.player2Name };
                // Initialize with the start word
                gameTurns.value = [{
                    Player: { ID: "", Name: "Start" },
                    Word: payload.startWord,
                    Timestamp: Date.now()
                }];
                break;

            case "word_submitted":
                // A word was played
                currentWord.value = payload.word;
                currentTurnId.value = payload.currentTurnId;
                gameTurns.value.push({
                    Player: { ID: payload.playerId, Name: payload.playerName },
                    Word: payload.word,
                    Timestamp: Date.now()
                });
                break;

            case "game_ended":
            case "game_ended_timeout":
                // Game has ended
                gameStatus.value = "completed";
                winnerId.value = payload?.winnerId || null;
                winReason.value = payload?.reason || (type === "game_ended_timeout" ? "timeout" : null);
                break;

            case "error":
                // Handle error from server
                error.value = payload?.message || "Unknown error";
                console.error("Game error:", payload?.errorType, payload?.message);
                break;

            default:
                console.log("Unknown game update type:", type);
        }
    }

    const LeaveGame = () => {
        useSocketStore().sendSocketMessage(JSON.stringify({
            action: "leave_game",
            gameId: gameId.value
        }));
        $reset();
    }


    return { 
        gameId, gameType, gameStatus, currentWord, currentTurnId, 
        player1, player2, gameTurns, winnerId, winReason, joinCode, error,
        FindGame, CreatePrivateGame, JoinPrivateGame, SendWord, HandleGameUpdate, LeaveGame, $reset
    }

});