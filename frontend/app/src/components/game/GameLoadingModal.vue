<script setup lang="ts">
import { computed, onUnmounted, ref, watch } from "vue"
import Modal from '@/components/modal/Modal.vue'
import { ClipboardDocumentIcon, CheckIcon, UserGroupIcon } from '@heroicons/vue/24/outline'

const props = defineProps<{
    showModal: boolean
    joinCode: string
}>()

const emit = defineEmits(['cancel'])

const dots = ref('')
const copied = ref(false)
let intervalId = -1

const animateDots = () => {
    let count = 0
    intervalId = setInterval(() => {
        count = (count + 1) % 4
        dots.value = '.'.repeat(count)
    }, 500)
}

const stopAnimateDots = () => {
    if (intervalId > -1) {
        clearInterval(intervalId)
        dots.value = ''
        intervalId = -1
    }
}

watch(() => props.showModal, (newV: boolean) => {
    if (newV) {
        stopAnimateDots()
        animateDots()
        copied.value = false
    } else {
        stopAnimateDots()
    }
})

onUnmounted(() => {
    stopAnimateDots()
})

const isPrivateGame = computed(() => {
    return props.joinCode?.length > 0
})

const modalTitle = computed(() => {
    return isPrivateGame.value ? "Private Game" : "Finding Match"
})

const copyJoinCode = async () => {
    if (!props.joinCode) return
    try {
        await navigator.clipboard.writeText(props.joinCode)
        copied.value = true
        setTimeout(() => {
            copied.value = false
        }, 2000)
    } catch (err) {
        console.error('Failed to copy:', err)
    }
}

const handleCancel = () => {
    emit('cancel')
}
</script>

<template>
    <Modal 
        id="gameLoadingModal" 
        :title="modalTitle" 
        :active="showModal" 
        :showCloseBtn="false"
    >
        <template #content>
            <div class="loading-content">
                <!-- Animated Icon -->
                <div class="loading-animation">
                    <div class="pulse-ring"></div>
                    <div class="pulse-ring delay-1"></div>
                    <div class="pulse-ring delay-2"></div>
                    <UserGroupIcon class="loading-icon" />
                </div>

                <!-- Join Code (for private games) -->
                <div v-if="isPrivateGame" class="join-code-section">
                    <span class="code-label">Share this code with your friend</span>
                    <div class="code-display">
                        <span class="code-text">{{ joinCode }}</span>
                        <button class="copy-btn" @click="copyJoinCode" :class="{ copied }">
                            <CheckIcon v-if="copied" class="copy-icon" />
                            <ClipboardDocumentIcon v-else class="copy-icon" />
                        </button>
                    </div>
                    <span v-if="copied" class="copied-text">Copied!</span>
                </div>

                <!-- Waiting Message -->
                <div class="waiting-message">
                    <span class="waiting-text">
                        {{ isPrivateGame ? 'Waiting for player to join' : 'Looking for an opponent' }}
                        <span class="dots">{{ dots }}</span>
                    </span>
                </div>
            </div>
        </template>
        <template #footerContent>
            <button class="cancel-btn" @click="handleCancel">
                Cancel
            </button>
        </template>
    </Modal>
</template>

<style scoped>
.loading-content {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 1.5rem;
    padding: 1rem 0;
}

/* Loading Animation */
.loading-animation {
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 80px;
    height: 80px;
}

.pulse-ring {
    position: absolute;
    width: 100%;
    height: 100%;
    border: 2px solid rgba(99, 102, 241, 0.3);
    border-radius: 50%;
    animation: pulse-out 2s ease-out infinite;
}

.pulse-ring.delay-1 {
    animation-delay: 0.5s;
}

.pulse-ring.delay-2 {
    animation-delay: 1s;
}

@keyframes pulse-out {
    0% {
        transform: scale(0.5);
        opacity: 1;
    }
    100% {
        transform: scale(1.5);
        opacity: 0;
    }
}

.loading-icon {
    width: 2.5rem;
    height: 2.5rem;
    color: #6366f1;
    z-index: 1;
}

/* Join Code Section */
.join-code-section {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 0.75rem;
}

.code-label {
    font-size: 0.875rem;
    color: #9ca3af;
}

.code-display {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.75rem 1rem;
    background: rgba(99, 102, 241, 0.1);
    border: 1px solid rgba(99, 102, 241, 0.3);
    border-radius: 0.75rem;
}

.code-text {
    font-family: monospace;
    font-size: 1.5rem;
    font-weight: 700;
    letter-spacing: 0.15em;
    color: #a5b4fc;
}

.copy-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 2rem;
    height: 2rem;
    background: rgba(255, 255, 255, 0.1);
    border: none;
    border-radius: 0.375rem;
    cursor: pointer;
    transition: all 0.2s ease;
}

.copy-btn:hover {
    background: rgba(255, 255, 255, 0.2);
}

.copy-btn.copied {
    background: rgba(34, 197, 94, 0.2);
}

.copy-icon {
    width: 1rem;
    height: 1rem;
    color: #9ca3af;
}

.copy-btn.copied .copy-icon {
    color: #22c55e;
}

.copied-text {
    font-size: 0.75rem;
    color: #22c55e;
    font-weight: 500;
}

/* Waiting Message */
.waiting-message {
    text-align: center;
}

.waiting-text {
    font-family: monospace;
    font-size: 0.9rem;
    color: #6b7280;
}

.dots {
    display: inline-block;
    width: 1.5rem;
    text-align: left;
}

/* Cancel Button */
.cancel-btn {
    flex: 1;
    padding: 0.75rem 1.25rem;
    background: rgba(255, 255, 255, 0.05);
    border: 1px solid rgba(255, 255, 255, 0.15);
    border-radius: 0.75rem;
    color: #9ca3af;
    font-family: monospace;
    font-size: 0.95rem;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.2s ease;
}

.cancel-btn:hover {
    background: rgba(239, 68, 68, 0.1);
    border-color: rgba(239, 68, 68, 0.3);
    color: #fca5a5;
}
</style>