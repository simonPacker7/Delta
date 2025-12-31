import { defineStore } from "pinia";
import { ref } from "vue";
import { useGameStore } from "./game";
    
export const useSocketStore = defineStore("socket", () => {
    const socket = ref<WebSocket | null>()

    const $reset = () => {
        socket.value = null
    }

    const connectSocket = () => {
        if (socket.value) {
            return
        }
        socket.value = new WebSocket('ws://localhost/ws')
        socket.value.onmessage = onSocketMessage
    }

    const closeSocket = () => {
        if (socket.value) {
            socket.value.close()
            socket.value = null
        }  
    }

    const onSocketMessage = (event: MessageEvent<any>) => {
        const message = event.data as string
        console.log(`wsMessage: ${message}`)

        if (message == null || message == "") {
            return
        }

        useGameStore().HandleGameUpdate(message);

        // Handle different message types here
    }

    const sendSocketMessage = (message: string) => {
        if (socket.value && socket.value.readyState === WebSocket.OPEN) {
            socket.value.send(message)
        }
    }

    // Might need to have specific methods for joining / leaving a game (Or leave that up to game store)

    return { socket, connectSocket, closeSocket, sendSocketMessage }
});