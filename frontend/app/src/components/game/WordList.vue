<script setup lang="ts">
import { ref, onUnmounted, computed, watch, nextTick } from 'vue'
import type { GameTurn } from '@/stores/game'
import WordItem from "./WordItem.vue"

const props = defineProps<{
    gameMoves?: GameTurn[]
    isGameActive: boolean
    currentPlayerId?: string | null
}>()

const emit = defineEmits(['timeExpired'])

const secondsLeft = ref(0)
const TURN_TIME_LIMIT = 99
let secondsLeftInterval: number | null = null
const scrollContainerRef = ref<HTMLElement | null>(null)

const clearSecondsInterval = () => {
    if (secondsLeftInterval != null) {
        clearInterval(secondsLeftInterval)
        secondsLeftInterval = null
    }
}

const restartSecondsInterval = () => {
    clearSecondsInterval()
    secondsLeft.value = TURN_TIME_LIMIT
    secondsLeftInterval = setInterval(() => {
        if (secondsLeft.value > 0) {
            secondsLeft.value--
        } else {
            clearSecondsInterval()
            emit("timeExpired")
        }
    }, 1000)
}

const scrollToBottom = () => {
    nextTick(() => {
        if (scrollContainerRef.value) {
            scrollContainerRef.value.scrollTo({
                top: scrollContainerRef.value.scrollHeight,
                behavior: 'smooth'
            })
        }
    })
}

// Watch for new words and scroll to bottom
watch(() => props.gameMoves?.length, (newLength, oldLength) => {
    if (newLength && (!oldLength || newLength > oldLength)) {
        if (props.isGameActive) {
            restartSecondsInterval()
        }
        scrollToBottom()
    }
})

// Also handle game becoming inactive
watch(() => props.isGameActive, (active) => {
    if (!active) {
        clearSecondsInterval()
    }
})

const isNewWord = (index: number) => {
    return props.gameMoves != null
        && index == props.gameMoves.length - 1
        && props.isGameActive
}

const isCurrentPlayerMove = (playerId: string) => {
    return playerId === props.currentPlayerId
}

const timerPercentage = computed(() => {
    return (secondsLeft.value / TURN_TIME_LIMIT) * 100
})

const timerColor = computed(() => {
    if (secondsLeft.value > 30) return '#22c55e'
    if (secondsLeft.value > 10) return '#eab308'
    return '#ef4444'
})

onUnmounted(() => {
    clearSecondsInterval()
})
</script>

<template>
    <div class="word-list-container">
        <!-- Timer Bar (only visible during active game) -->
        <div v-if="isGameActive && gameMoves && gameMoves.length > 0" class="timer-bar">
            <div class="timer-track">
                <div 
                    class="timer-fill"
                    :style="{ 
                        width: timerPercentage + '%',
                        backgroundColor: timerColor
                    }"
                ></div>
            </div>
            <span class="timer-text" :style="{ color: timerColor }">
                {{ secondsLeft }}s
            </span>
        </div>

        <!-- Words List -->
        <div ref="scrollContainerRef" class="words-scroll">
            <TransitionGroup name="word" tag="div" class="words-list">
                <div
                    v-for="({ Player, Word }, index) in gameMoves"
                    :key="index"
                    class="word-row"
                    :class="{ 
                        'is-new': isNewWord(index),
                        'is-current-player': isCurrentPlayerMove(Player.ID)
                    }"
                >
                    <!-- Player indicator -->
                    <div class="player-indicator" :class="{ 'self': isCurrentPlayerMove(Player.ID) }">
                        <span class="player-label">{{ Player.Name }}</span>
                    </div>

                    <!-- Word display -->
                    <WordItem :word="Word" :highlight="isNewWord(index)" />
                </div>
            </TransitionGroup>

            <!-- Empty state -->
            <div v-if="!gameMoves || gameMoves.length === 0" class="empty-state">
                <span class="empty-text">Waiting for first word...</span>
            </div>
        </div>
    </div>
</template>

<style scoped>
.word-list-container {
    display: flex;
    flex-direction: column;
    flex: 1;
    min-height: 0;
    background: rgba(255, 255, 255, 0.02);
    border: 1px solid rgba(255, 255, 255, 0.08);
    border-radius: 1rem;
    overflow: hidden;
}

/* Timer Bar */
.timer-bar {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    padding: 0.75rem 1rem;
    background: rgba(0, 0, 0, 0.2);
    border-bottom: 1px solid rgba(255, 255, 255, 0.05);
}

.timer-track {
    flex: 1;
    height: 6px;
    background: rgba(255, 255, 255, 0.1);
    border-radius: 3px;
    overflow: hidden;
}

.timer-fill {
    height: 100%;
    border-radius: 3px;
    transition: width 1s linear, background-color 0.3s ease;
}

.timer-text {
    font-family: monospace;
    font-size: 0.875rem;
    font-weight: 600;
    min-width: 3rem;
    text-align: right;
}

/* Words scroll area */
.words-scroll {
    flex: 1;
    overflow-y: auto;
    overflow-x: hidden;
    padding: 0.5rem;
    display: flex;
    flex-direction: column;
}

/* Hide scrollbar but keep functionality */
.words-scroll::-webkit-scrollbar {
    width: 6px;
}

.words-scroll::-webkit-scrollbar-track {
    background: transparent;
}

.words-scroll::-webkit-scrollbar-thumb {
    background: rgba(255, 255, 255, 0.1);
    border-radius: 3px;
}

.words-scroll::-webkit-scrollbar-thumb:hover {
    background: rgba(255, 255, 255, 0.2);
}

/* Words list */
.words-list {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
    margin-top: auto;
}

/* Word row */
.word-row {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 0.375rem;
    padding: 0.75rem;
    background: rgba(255, 255, 255, 0.02);
    border: 1px solid rgba(255, 255, 255, 0.05);
    border-radius: 0.75rem;
    transition: all 0.3s ease;
}

.word-row.is-new {
    background: rgba(34, 197, 94, 0.08);
    border-color: rgba(34, 197, 94, 0.2);
}

.word-row.is-current-player {
    background: rgba(99, 102, 241, 0.08);
    border-color: rgba(99, 102, 241, 0.15);
}

.word-row.is-current-player.is-new {
    background: linear-gradient(135deg, rgba(99, 102, 241, 0.1) 0%, rgba(34, 197, 94, 0.1) 100%);
    border-color: rgba(34, 197, 94, 0.25);
}

/* Player indicator */
.player-indicator {
    display: flex;
    align-items: center;
    gap: 0.375rem;
}

.player-label {
    font-family: monospace;
    font-size: 0.7rem;
    font-weight: 500;
    text-transform: uppercase;
    letter-spacing: 0.05em;
    color: #6b7280;
}

.player-indicator.self .player-label {
    color: #a5b4fc;
}

/* Empty state */
.empty-state {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 2rem;
}

.empty-text {
    font-family: monospace;
    font-size: 0.9rem;
    color: #4b5563;
}

/* Word transition animations */
.word-enter-active {
    transition: all 0.4s ease;
}

.word-leave-active {
    transition: all 0.2s ease;
}

.word-enter-from {
    opacity: 0;
    transform: translateY(20px) scale(0.95);
}

.word-leave-to {
    opacity: 0;
}

/* Responsive */
@media (max-width: 480px) {
    .timer-bar {
        padding: 0.625rem 0.75rem;
    }

    .timer-text {
        font-size: 0.8rem;
        min-width: 2.5rem;
    }

    .word-row {
        padding: 0.625rem;
    }

    .player-label {
        font-size: 0.65rem;
    }
}
</style>