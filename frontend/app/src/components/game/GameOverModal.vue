<script setup lang="ts">
import { computed } from 'vue'
import Modal from '../modal/Modal.vue'
import { TrophyIcon, XCircleIcon, ClockIcon, NoSymbolIcon } from '@heroicons/vue/24/outline'

const props = defineProps<{
    winningReason?: string
    isWinningPlayer?: boolean
    showModal: boolean
}>()

const emit = defineEmits(['exitGame'])

const onLeaveGame = () => {
    emit('exitGame')
}

const reasonMessage = computed(() => {
    switch (props.winningReason) {
        case "no possible words":
            return "No more valid words"
        case "timelimit":
            return "Time ran out"
        case "forfeit":
            return "Opponent left the game"
        default:
            return "Game complete"
    }
})

const reasonIcon = computed(() => {
    switch (props.winningReason) {
        case "no possible words":
            return NoSymbolIcon
        case "timelimit":
            return ClockIcon
        default:
            return null
    }
})

const resultTitle = computed(() => {
    return props.isWinningPlayer ? "Victory!" : "Defeat"
})

const resultMessage = computed(() => {
    if (props.isWinningPlayer) {
        return "Congratulations! You outplayed your opponent."
    } else {
        return "Better luck next time! Keep practicing."
    }
})
</script>

<template>
    <Modal 
        id="gameOverModal" 
        title="Game Over" 
        :active="showModal" 
        :showCloseBtn="false"
    >
        <template #content>
            <div class="game-over-content">
                <!-- Result Icon -->
                <div class="result-icon-wrapper" :class="{ win: isWinningPlayer, loss: !isWinningPlayer }">
                    <TrophyIcon v-if="isWinningPlayer" class="result-icon" />
                    <XCircleIcon v-else class="result-icon" />
                </div>

                <!-- Result Title -->
                <h2 class="result-title" :class="{ win: isWinningPlayer, loss: !isWinningPlayer }">
                    {{ resultTitle }}
                </h2>

                <!-- Reason Badge -->
                <div class="reason-badge">
                    <component v-if="reasonIcon" :is="reasonIcon" class="reason-icon" />
                    <span>{{ reasonMessage }}</span>
                </div>

                <!-- Result Message -->
                <p class="result-message">{{ resultMessage }}</p>
            </div>
        </template>
        <template #footerContent>
            <button class="exit-btn" :class="{ win: isWinningPlayer }" @click="onLeaveGame">
                Continue
            </button>
        </template>
    </Modal>
</template>

<style scoped>
.game-over-content {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 1rem;
    padding: 1rem 0;
    text-align: center;
}

/* Result Icon */
.result-icon-wrapper {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 5rem;
    height: 5rem;
    border-radius: 50%;
    margin-bottom: 0.5rem;
}

.result-icon-wrapper.win {
    background: rgba(34, 197, 94, 0.15);
    animation: celebrate 0.6s ease;
}

.result-icon-wrapper.loss {
    background: rgba(239, 68, 68, 0.15);
}

@keyframes celebrate {
    0% {
        transform: scale(0.5);
        opacity: 0;
    }
    50% {
        transform: scale(1.1);
    }
    100% {
        transform: scale(1);
        opacity: 1;
    }
}

.result-icon {
    width: 2.5rem;
    height: 2.5rem;
}

.result-icon-wrapper.win .result-icon {
    color: #22c55e;
}

.result-icon-wrapper.loss .result-icon {
    color: #ef4444;
}

/* Result Title */
.result-title {
    font-family: monospace;
    font-size: 2rem;
    font-weight: 700;
    margin: 0;
}

.result-title.win {
    background: linear-gradient(135deg, #22c55e 0%, #10b981 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    background-clip: text;
}

.result-title.loss {
    color: #ef4444;
}

/* Reason Badge */
.reason-badge {
    display: inline-flex;
    align-items: center;
    gap: 0.375rem;
    padding: 0.5rem 1rem;
    background: rgba(255, 255, 255, 0.05);
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 2rem;
    font-family: monospace;
    font-size: 0.8rem;
    color: #9ca3af;
}

.reason-icon {
    width: 1rem;
    height: 1rem;
}

/* Result Message */
.result-message {
    font-size: 0.95rem;
    color: #6b7280;
    margin: 0;
    max-width: 280px;
    line-height: 1.5;
}

/* Exit Button */
.exit-btn {
    flex: 1;
    padding: 0.875rem 1.5rem;
    border: none;
    border-radius: 0.75rem;
    font-family: monospace;
    font-size: 1rem;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.2s ease;
    background: rgba(255, 255, 255, 0.1);
    color: #ffffff;
}

.exit-btn:hover {
    background: rgba(255, 255, 255, 0.15);
}

.exit-btn.win {
    background: linear-gradient(135deg, #22c55e 0%, #16a34a 100%);
}

.exit-btn.win:hover {
    box-shadow: 0 8px 20px rgba(34, 197, 94, 0.3);
    transform: translateY(-2px);
}
</style>