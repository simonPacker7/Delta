<script setup lang="ts">
import { computed } from 'vue'
import { useGameStore } from '@/stores/game'
import { useSessionStore } from '@/stores/session'
import { useRouter } from 'vue-router'

import InputBar from '@/components/game/InputBar.vue'
import WordList from '@/components/game/WordList.vue'
import GameOverModal from '@/components/game/GameOverModal.vue'
import GameLoadingModal from "@/components/game/GameLoadingModal.vue"

const gameStore = useGameStore()
const sessionStore = useSessionStore()
const router = useRouter()

const isGameActive = computed(() => {
    return gameStore.gameStatus !== 'completed'
})

const isTurn = computed(() => gameStore.currentTurnId == sessionStore.id)

const SendWord = (word: string) => {
    return gameStore.SendWord(word)
}

const isLoading = computed(() => {
    return gameStore.gameStatus == 'waiting' || gameStore.gameStatus == 'ready'
})

const isGameOver = computed(() => {
    return gameStore.gameStatus == 'completed'
})

const isWinningPlayer = computed(() => {
    return gameStore.winnerId == sessionStore.id
})

const winningReason = computed(() => {
    return gameStore.winReason
})

const onLeaveGame = () => {
    gameStore.LeaveGame()
    router.replace("/play")
}
</script>

<template>
    <div class="game-container">
        <WordList :game-moves="gameStore.gameTurns" :is-game-active="isGameActive" />
        <InputBar :is-turn="isTurn" :submit-word="SendWord" />
        <GameLoadingModal :show-modal="isLoading" :join-code="gameStore.joinCode" />
        <GameOverModal :show-modal="isGameOver" :is-winning-player="isWinningPlayer" :winning-reason="winningReason"
            @exit-game="onLeaveGame" />
    </div>
</template>

<style scoped>
.game-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    height: 80vh;
}
</style>