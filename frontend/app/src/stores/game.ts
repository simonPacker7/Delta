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
    const currentPlayerId = ref<string | null>(null);
    const gameTurns = ref<GameTurn[]>([]);
    const winningPlayerId = ref<string | null>(null);
    const winningReason = ref<string | null>(null);
    const joinCode = ref<string | null>(null);

    const $reset = () => {
        gameId.value = null;
        gameType.value = null;
        gameStatus.value = null;
        currentPlayerId.value = null;
        gameTurns.value = [];
        winningPlayerId.value = null;
        winningReason.value = null;
        joinCode.value = null;
    }

    const FindGame = async () => {
        try{
            const res = await axios.get('/api/game/find');
            if (!res || res.status != 200 || !res.data || !res.data.gameId) {
                $reset();
                return;
            }
            gameId.value = res.data.gameId;
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

    const HandleGameUpdate = (data: any) => {
        console.log("Handling game update:", data);
        const update = JSON.parse(data);
    }


    return { gameId, gameType, gameStatus, currentPlayerId, gameTurns, winningPlayerId, winningReason, joinCode, FindGame, SendWord, HandleGameUpdate  }

});