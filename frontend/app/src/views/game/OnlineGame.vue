<script setup lang="ts">
import { computed, ref, onMounted, onUnmounted } from 'vue'
import { useGameStore } from '@/stores/game'
import { useSessionStore } from '@/stores/session'
import { useRouter, onBeforeRouteLeave } from 'vue-router'

import InputBar from '@/components/game/InputBar.vue'
import WordList from '@/components/game/WordList.vue'
import GameOverModal from '@/components/game/GameOverModal.vue'
import GameLoadingModal from "@/components/game/GameLoadingModal.vue"
import GameHeader from '@/components/game/GameHeader.vue'
import ForfeitConfirmModal from '@/components/game/ForfeitConfirmModal.vue'

const gameStore = useGameStore()
const sessionStore = useSessionStore()
const router = useRouter()

// Forfeit modal state
const showForfeitModal = ref(false)
const isNavigationForfeit = ref(false)
const pendingNavigation = ref<string | null>(null)
const isForfeiting = ref(false)

const isGameActive = computed(() => {
    return gameStore.gameStatus !== 'completed'
})

const isTurn = computed(() => gameStore.currentTurnId == sessionStore.id)

const currentPlayerId = computed(() => sessionStore.id)

const SendWord = (word: string) => {
    gameStore.SendWord(word)
}

const gameError = computed(() => gameStore.error)

const clearError = () => {
    gameStore.clearError()
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

const onCancelMatchmaking = async () => {
    await gameStore.CancelMatchmaking()
    router.replace("/play")
}

// Forfeit functionality
const onForfeitClick = () => {
    isNavigationForfeit.value = false
    showForfeitModal.value = true
}

const onConfirmForfeit = () => {
    showForfeitModal.value = false
    isForfeiting.value = true
    gameStore.ForfeitGame()
    
    if (pendingNavigation.value) {
        router.push(pendingNavigation.value)
        pendingNavigation.value = null
    }
}

const onCancelForfeit = () => {
    showForfeitModal.value = false
    pendingNavigation.value = null
}

// Navigation guard for forfeit warning
onBeforeRouteLeave((to) => {
    // Skip guard if we're already in the process of forfeiting
    if (isForfeiting.value) {
        return true
    }
    
    // Only show warning if game is active (not waiting/ready/completed)
    if (gameStore.gameStatus === 'active') {
        isNavigationForfeit.value = true
        pendingNavigation.value = to.fullPath
        showForfeitModal.value = true
        return false
    }
    return true
})

// Handle browser back/refresh during active game
const handleBeforeUnload = (e: BeforeUnloadEvent) => {
    if (gameStore.gameStatus === 'active') {
        e.preventDefault()
        e.returnValue = ''
    }
}

onMounted(() => {
    window.addEventListener('beforeunload', handleBeforeUnload)
})

onUnmounted(() => {
    window.removeEventListener('beforeunload', handleBeforeUnload)
})

// Player info for the header
const player1 = computed(() => gameStore.player1)
const player2 = computed(() => gameStore.player2)
</script>

<template>
    <div class="game-page">
        <div class="game-container">
            <!-- Game Header with player info -->
            <GameHeader 
                :player1="player1"
                :player2="player2"
                :current-turn-id="gameStore.currentTurnId"
                :current-player-id="currentPlayerId"
                :is-game-active="isGameActive"
            />
            
            <!-- Main Game Area -->
            <div class="game-area">
                <WordList 
                    :game-moves="gameStore.gameTurns" 
                    :is-game-active="isGameActive"
                    :current-player-id="currentPlayerId"
                />
            </div>
            
            <!-- Input Area -->
            <InputBar 
                :is-turn="isTurn" 
                :submit-word="SendWord"
                :is-game-active="isGameActive"
                :error="gameError"
                :clear-error="clearError"
            />

            <!-- Forfeit Button -->
            <button 
                v-if="gameStore.gameStatus === 'active'" 
                class="forfeit-btn"
                @click="onForfeitClick"
            >
                Forfeit
            </button>
        </div>
        
        <!-- Modals -->
        <GameLoadingModal 
            :show-modal="isLoading" 
            :join-code="gameStore.joinCode"
            @cancel="onCancelMatchmaking"
        />
        <GameOverModal 
            :show-modal="isGameOver" 
            :is-winning-player="isWinningPlayer" 
            :winning-reason="winningReason"
            @exit-game="onLeaveGame" 
        />
        <ForfeitConfirmModal
            :show-modal="showForfeitModal"
            :is-navigation-warning="isNavigationForfeit"
            @confirm="onConfirmForfeit"
            @cancel="onCancelForfeit"
        />

        <!-- Background decoration -->
        <div class="bg-decoration">
            <div class="gradient-orb orb-1"></div>
            <div class="gradient-orb orb-2"></div>
        </div>
    </div>
</template>

<style scoped>
.game-page {
    position: relative;
    min-height: calc(100vh - 70px);
    display: flex;
    justify-content: center;
    padding: 1.5rem;
}

.game-container {
    position: relative;
    z-index: 10;
    display: flex;
    flex-direction: column;
    width: 100%;
    max-width: 600px;
    height: calc(100vh - 70px - 3rem);
}

.game-area {
    flex: 1;
    display: flex;
    flex-direction: column;
    min-height: 0;
    margin-bottom: 1rem;
}

/* Forfeit Button */
.forfeit-btn {
    font-family: monospace;
    font-size: 0.85rem;
    font-weight: 600;
    padding: 0.5rem 1rem;
    margin-top: 0.75rem;
    border-radius: 0.5rem;
    border: 1px solid rgba(239, 68, 68, 0.3);
    background: rgba(239, 68, 68, 0.1);
    color: #ef4444;
    cursor: pointer;
    transition: all 0.2s ease;
    align-self: center;
}

.forfeit-btn:hover {
    background: rgba(239, 68, 68, 0.2);
    border-color: rgba(239, 68, 68, 0.5);
}

/* Background decoration */
.bg-decoration {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    pointer-events: none;
    z-index: 1;
    overflow: hidden;
}

.gradient-orb {
    position: absolute;
    border-radius: 50%;
    filter: blur(100px);
    opacity: 0.2;
}

.orb-1 {
    width: 400px;
    height: 400px;
    background: radial-gradient(circle, #6366f1 0%, transparent 70%);
    top: -100px;
    right: -100px;
}

.orb-2 {
    width: 350px;
    height: 350px;
    background: radial-gradient(circle, #22c55e 0%, transparent 70%);
    bottom: -100px;
    left: -100px;
}

/* Responsive */
@media (max-width: 640px) {
    .game-page {
        padding: 1rem;
    }

    .game-container {
        height: calc(100vh - 60px - 2rem);
    }
}

@media (max-width: 480px) {
    .game-page {
        padding: 0.75rem;
    }
}
</style>