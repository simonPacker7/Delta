<script setup lang="ts">
import { ref, watch } from 'vue'
import Modal from '@/components/modal/Modal.vue'
import { KeyIcon } from '@heroicons/vue/24/outline'

const props = defineProps({
    showModal: { type: Boolean, required: true },
    joinGameFunc: { type: Function, required: true },
})
const emit = defineEmits(['cancel', 'joinedGame'])

const joinCode = ref('')
const isLoading = ref(false)
const errorMessage = ref('')
const inputRef = ref<HTMLInputElement | null>(null)

watch(() => props.showModal, (newVal) => {
    joinCode.value = ''
    errorMessage.value = ''
    if (newVal) {
        // Focus input when modal opens
        setTimeout(() => inputRef.value?.focus(), 100)
    }
})

const join = async () => {
    if (!joinCode.value.trim()) {
        errorMessage.value = 'Please enter a game code'
        return
    }
    
    isLoading.value = true
    errorMessage.value = ''
    
    try {
        const error = await props.joinGameFunc(joinCode.value)
        if (error != null) {
            errorMessage.value = error
            return
        }
        emit('joinedGame')
    } finally {
        isLoading.value = false
    }
}

const handleCancel = () => {
    emit('cancel')
}
</script>

<template>
    <Modal 
        id="joinGameModal" 
        title="Join Private Game" 
        :active="showModal" 
        :showCloseBtn="true"
        @close="handleCancel"
    >
        <template #content>
            <form @submit.prevent="join" class="join-form">
                <div class="form-group">
                    <label for="joinCode" class="form-label">Game Code</label>
                    <div class="input-wrapper">
                        <KeyIcon class="input-icon" />
                        <input 
                            ref="inputRef"
                            v-model="joinCode" 
                            id="joinCode"
                            type="text"
                            class="form-input"
                            placeholder="Enter the code from your friend"
                            :disabled="isLoading"
                            autocomplete="off"
                        />
                    </div>
                </div>
                
                <div v-if="errorMessage" class="error-message">
                    {{ errorMessage }}
                </div>
            </form>
        </template>
        <template #footerContent>
            <button 
                type="button" 
                class="modal-btn secondary" 
                @click="handleCancel"
                :disabled="isLoading"
            >
                Cancel
            </button>
            <button 
                type="button" 
                class="modal-btn primary" 
                @click="join"
                :disabled="isLoading || !joinCode.trim()"
            >
                <span v-if="isLoading" class="loading-spinner"></span>
                <span>{{ isLoading ? 'Joining...' : 'Join Game' }}</span>
            </button>
        </template>
    </Modal>
</template>

<style scoped>
.join-form {
    display: flex;
    flex-direction: column;
    gap: 1rem;
}

.form-group {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
}

.form-label {
    font-size: 0.875rem;
    font-weight: 500;
    color: #d1d5db;
    font-family: monospace;
}

.input-wrapper {
    position: relative;
    display: flex;
    align-items: center;
}

.input-icon {
    position: absolute;
    left: 1rem;
    width: 1.25rem;
    height: 1.25rem;
    color: #6b7280;
    pointer-events: none;
}

.form-input {
    width: 100%;
    padding: 0.875rem 1rem 0.875rem 3rem;
    background: rgba(255, 255, 255, 0.05);
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 0.75rem;
    color: #ffffff;
    font-size: 1rem;
    font-family: monospace;
    transition: all 0.2s ease;
}

.form-input::placeholder {
    color: #6b7280;
}

.form-input:focus {
    outline: none;
    border-color: #6366f1;
    box-shadow: 0 0 0 3px rgba(99, 102, 241, 0.2);
    background: rgba(255, 255, 255, 0.08);
}

.form-input:disabled {
    opacity: 0.6;
    cursor: not-allowed;
}

.error-message {
    background: rgba(239, 68, 68, 0.1);
    border: 1px solid rgba(239, 68, 68, 0.3);
    color: #fca5a5;
    padding: 0.75rem 1rem;
    border-radius: 0.5rem;
    font-size: 0.875rem;
    text-align: center;
}

.modal-btn {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
    padding: 0.75rem 1.25rem;
    border-radius: 0.75rem;
    font-size: 0.95rem;
    font-weight: 600;
    font-family: monospace;
    cursor: pointer;
    transition: all 0.2s ease;
    border: none;
}

.modal-btn.primary {
    background: linear-gradient(135deg, #6366f1 0%, #8b5cf6 100%);
    color: #ffffff;
}

.modal-btn.primary:hover:not(:disabled) {
    box-shadow: 0 8px 20px rgba(99, 102, 241, 0.4);
}

.modal-btn.primary:disabled {
    opacity: 0.5;
    cursor: not-allowed;
}

.modal-btn.secondary {
    background: rgba(255, 255, 255, 0.05);
    border: 1px solid rgba(255, 255, 255, 0.15);
    color: #9ca3af;
}

.modal-btn.secondary:hover:not(:disabled) {
    background: rgba(255, 255, 255, 0.1);
    color: #ffffff;
}

.loading-spinner {
    width: 1rem;
    height: 1rem;
    border: 2px solid rgba(255, 255, 255, 0.3);
    border-top-color: #ffffff;
    border-radius: 50%;
    animation: spin 0.8s linear infinite;
}

@keyframes spin {
    to {
        transform: rotate(360deg);
    }
}
</style>