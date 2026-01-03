<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { PaperAirplaneIcon } from '@heroicons/vue/24/solid'

const props = defineProps<{
    isTurn: boolean
    submitWord: (word: string) => void
    isGameActive?: boolean
    error?: string | null
    clearError?: () => void
}>()

const curText = ref('')
const isSubmitting = ref(false)
const localErrorMessage = ref('')
const inputRef = ref<HTMLInputElement | null>(null)

// Computed to use either prop error or local error
const errorMessage = computed(() => props.error || localErrorMessage.value)

// Error message mapping for user-friendly display
const errorDisplayMessages: Record<string, string> = {
    'word_already_played': 'Word already played!',
    'invalid_move': 'Invalid word change!',
    'not_your_turn': "Not your turn!",
    'not_in_game': 'Not in game!',
    'game_not_found': 'Game not found!',
}

const displayError = computed(() => {
    const err = errorMessage.value
    if (!err) return ''
    return errorDisplayMessages[err] || err
})

// Watch for error changes from props and auto-clear after delay
watch(() => props.error, (newVal) => {
    if (newVal) {
        setTimeout(() => {
            props.clearError?.()
        }, 3000)
    }
})

// Focus input when it becomes the player's turn
watch(() => props.isTurn, (newVal) => {
    if (newVal && inputRef.value) {
        inputRef.value.focus()
    }
})

const turnStatus = computed(() => {
    if (!props.isGameActive) return 'Game Over'
    return props.isTurn ? "Your Turn" : "Opponent's Turn"
})

const statusClass = computed(() => ({
    'your-turn': props.isTurn && props.isGameActive,
    'opponent-turn': !props.isTurn && props.isGameActive,
    'game-over': !props.isGameActive
}))

const canSubmit = computed(() => {
    return props.isTurn && 
           curText.value.trim().length === 4 && 
           !isSubmitting.value &&
           props.isGameActive !== false
})

const SendWord = () => {
    if (!canSubmit.value) return

    // Clear any previous errors
    localErrorMessage.value = ''
    props.clearError?.()

    // Send the word - errors will come back via props.error
    props.submitWord(curText.value.toUpperCase())
    curText.value = ''
}

const handleKeydown = (e: KeyboardEvent) => {
    if (e.key === 'Enter' && canSubmit.value) {
        SendWord()
    }
}

// Only allow letters
const handleInput = (e: Event) => {
    const input = e.target as HTMLInputElement
    curText.value = input.value.replace(/[^a-zA-Z]/g, '').slice(0, 4)
}
</script>

<template>
    <div class="input-bar">
        <!-- Turn Status -->
        <div class="status-indicator" :class="statusClass">
            <div class="status-dot"></div>
            <span class="status-text">{{ turnStatus }}</span>
        </div>

        <!-- Input Area -->
        <div class="input-area">
            <div class="input-wrapper" :class="{ 'has-error': displayError, 'disabled': !isTurn }">
                <input
                    ref="inputRef"
                    :value="curText"
                    @input="handleInput"
                    @keydown="handleKeydown"
                    class="word-input"
                    type="text"
                    maxlength="4"
                    placeholder="____"
                    :disabled="!isTurn || isGameActive === false"
                    autocomplete="off"
                    autocapitalize="off"
                    spellcheck="false"
                />
                <div class="letter-count">{{ curText.length }}/4</div>
            </div>

            <button 
                class="submit-btn"
                :class="{ 'ready': canSubmit }"
                :disabled="!canSubmit"
                @click="SendWord"
            >
                <PaperAirplaneIcon v-if="!isSubmitting" class="submit-icon" />
                <span v-else class="loading-spinner"></span>
            </button>
        </div>

        <!-- Error Message -->
        <Transition name="fade">
            <div v-if="displayError" class="error-toast">
                {{ displayError }}
            </div>
        </Transition>
    </div>
</template>

<style scoped>
.input-bar {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 0.75rem;
    padding: 1rem;
    background: rgba(255, 255, 255, 0.03);
    border: 1px solid rgba(255, 255, 255, 0.08);
    border-radius: 1rem;
}

/* Status Indicator */
.status-indicator {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.375rem 0.875rem;
    background: rgba(255, 255, 255, 0.05);
    border-radius: 2rem;
}

.status-dot {
    width: 0.5rem;
    height: 0.5rem;
    border-radius: 50%;
    background: #6b7280;
    transition: all 0.3s ease;
}

.status-indicator.your-turn .status-dot {
    background: #22c55e;
    box-shadow: 0 0 8px rgba(34, 197, 94, 0.5);
    animation: pulse 2s ease-in-out infinite;
}

.status-indicator.opponent-turn .status-dot {
    background: #f59e0b;
}

.status-indicator.game-over .status-dot {
    background: #6b7280;
}

@keyframes pulse {
    0%, 100% {
        opacity: 1;
    }
    50% {
        opacity: 0.5;
    }
}

.status-text {
    font-family: monospace;
    font-size: 0.8rem;
    font-weight: 600;
    text-transform: uppercase;
    letter-spacing: 0.05em;
    color: #9ca3af;
}

.status-indicator.your-turn .status-text {
    color: #22c55e;
}

.status-indicator.opponent-turn .status-text {
    color: #f59e0b;
}

/* Input Area */
.input-area {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    width: 100%;
    max-width: 280px;
}

.input-wrapper {
    position: relative;
    flex: 1;
}

.word-input {
    width: 100%;
    padding: 0.875rem 1rem;
    background: rgba(255, 255, 255, 0.05);
    border: 2px solid rgba(255, 255, 255, 0.1);
    border-radius: 0.75rem;
    color: #ffffff;
    font-family: monospace;
    font-size: 1.5rem;
    font-weight: 700;
    text-align: center;
    text-transform: uppercase;
    letter-spacing: 0.3em;
    transition: all 0.2s ease;
}

.word-input::placeholder {
    color: #4b5563;
    letter-spacing: 0.5em;
}

.word-input:focus {
    outline: none;
    border-color: #6366f1;
    box-shadow: 0 0 0 3px rgba(99, 102, 241, 0.2);
    background: rgba(255, 255, 255, 0.08);
}

.word-input:disabled {
    opacity: 0.5;
    cursor: not-allowed;
}

.input-wrapper.disabled .word-input {
    border-color: rgba(255, 255, 255, 0.05);
}

.input-wrapper.has-error .word-input {
    border-color: #ef4444;
    animation: shake 0.4s ease;
}

@keyframes shake {
    0%, 100% { transform: translateX(0); }
    20%, 60% { transform: translateX(-5px); }
    40%, 80% { transform: translateX(5px); }
}

.letter-count {
    position: absolute;
    right: 0.75rem;
    bottom: -1.5rem;
    font-family: monospace;
    font-size: 0.7rem;
    color: #6b7280;
}

/* Submit Button */
.submit-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 3.5rem;
    height: 3.5rem;
    background: rgba(255, 255, 255, 0.05);
    border: 2px solid rgba(255, 255, 255, 0.1);
    border-radius: 0.75rem;
    cursor: pointer;
    transition: all 0.2s ease;
    flex-shrink: 0;
}

.submit-btn:disabled {
    opacity: 0.4;
    cursor: not-allowed;
}

.submit-btn.ready {
    background: linear-gradient(135deg, #22c55e 0%, #16a34a 100%);
    border-color: transparent;
}

.submit-btn.ready:hover {
    transform: scale(1.05);
    box-shadow: 0 8px 20px rgba(34, 197, 94, 0.3);
}

.submit-btn.ready:active {
    transform: scale(0.98);
}

.submit-icon {
    width: 1.5rem;
    height: 1.5rem;
    color: #6b7280;
    transform: rotate(90deg);
    transition: color 0.2s ease;
}

.submit-btn.ready .submit-icon {
    color: #ffffff;
}

.loading-spinner {
    width: 1.25rem;
    height: 1.25rem;
    border: 2px solid rgba(255, 255, 255, 0.3);
    border-top-color: #ffffff;
    border-radius: 50%;
    animation: spin 0.8s linear infinite;
}

@keyframes spin {
    to { transform: rotate(360deg); }
}

/* Error Toast */
.error-toast {
    position: absolute;
    bottom: 5.5rem;
    left: 50%;
    transform: translateX(-50%);
    padding: 0.5rem 1rem;
    background: rgba(239, 68, 68, 0.9);
    color: #ffffff;
    font-family: monospace;
    font-size: 0.8rem;
    font-weight: 500;
    border-radius: 0.5rem;
    white-space: nowrap;
}

.fade-enter-active,
.fade-leave-active {
    transition: opacity 0.2s ease, transform 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
    opacity: 0;
    transform: translateX(-50%) translateY(-5px);
}

/* Responsive */
@media (max-width: 480px) {
    .input-bar {
        padding: 0.875rem;
    }

    .input-area {
        max-width: 260px;
    }

    .word-input {
        padding: 0.75rem;
        font-size: 1.25rem;
        letter-spacing: 0.25em;
    }

    .submit-btn {
        width: 3rem;
        height: 3rem;
    }

    .submit-icon {
        width: 1.25rem;
        height: 1.25rem;
    }

    .status-text {
        font-size: 0.7rem;
    }
}
</style>