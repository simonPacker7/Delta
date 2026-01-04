<script setup lang="ts">
import Modal from '@/components/modal/Modal.vue'

defineProps<{
    showModal: boolean
    isNavigationWarning?: boolean
}>()

const emit = defineEmits<{
    confirm: []
    cancel: []
}>()
</script>

<template>
    <Modal 
        id="forfeit-confirm-modal" 
        :title="isNavigationWarning ? 'Leave Game?' : 'Forfeit Game?'" 
        :active="showModal"
        :show-close-btn="false"
    >
        <template #body>
            <div class="forfeit-content">
                <div class="warning-icon">⚠️</div>
                <p class="forfeit-message" v-if="isNavigationWarning">
                    Leaving now will forfeit the game and your opponent will win.
                </p>
                <p class="forfeit-message" v-else>
                    Are you sure you want to forfeit? Your opponent will be declared the winner.
                </p>
            </div>
        </template>
        <template #footer>
            <div class="forfeit-actions">
                <button class="cancel-btn" @click="emit('cancel')">
                    {{ isNavigationWarning ? 'Stay in Game' : 'Cancel' }}
                </button>
                <button class="confirm-btn" @click="emit('confirm')">
                    {{ isNavigationWarning ? 'Leave & Forfeit' : 'Forfeit' }}
                </button>
            </div>
        </template>
    </Modal>
</template>

<style scoped>
.forfeit-content {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 1.5rem;
    text-align: center;
}

.warning-icon {
    font-size: 3rem;
    margin-bottom: 1rem;
}

.forfeit-message {
    font-family: monospace;
    font-size: 1rem;
    color: rgba(255, 255, 255, 0.8);
    line-height: 1.5;
    margin: 0;
}

.forfeit-actions {
    display: flex;
    gap: 1rem;
    padding: 1rem 1.5rem 1.5rem;
    justify-content: center;
}

.cancel-btn,
.confirm-btn {
    font-family: monospace;
    font-size: 0.95rem;
    font-weight: 600;
    padding: 0.75rem 1.5rem;
    border-radius: 0.75rem;
    border: none;
    cursor: pointer;
    transition: all 0.2s ease;
}

.cancel-btn {
    background: rgba(255, 255, 255, 0.1);
    color: rgba(255, 255, 255, 0.8);
    border: 1px solid rgba(255, 255, 255, 0.2);
}

.cancel-btn:hover {
    background: rgba(255, 255, 255, 0.15);
    border-color: rgba(255, 255, 255, 0.3);
}

.confirm-btn {
    background: linear-gradient(135deg, #ef4444, #dc2626);
    color: white;
}

.confirm-btn:hover {
    background: linear-gradient(135deg, #dc2626, #b91c1c);
    transform: translateY(-1px);
}
</style>
