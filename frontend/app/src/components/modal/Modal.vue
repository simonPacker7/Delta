<script setup lang="ts">
import { XMarkIcon } from '@heroicons/vue/24/outline'

const props = defineProps({
    id: String,
    title: String,
    active: Boolean,
    showCloseBtn: {
        type: Boolean,
        default: true,
    },
})

const emit = defineEmits(['close'])
</script>

<template>
    <Teleport to="body">
        <Transition name="modal">
            <div v-if="active" :id="props.id" class="modal-overlay" role="dialog" aria-modal="true">
                <div class="modal-backdrop" @click="emit('close')"></div>
                <div class="modal-container">
                    <div class="modal-content">
                        <slot name="header">
                            <header class="modal-header">
                                <h2 class="modal-title">{{ props.title }}</h2>
                                <button 
                                    v-if="showCloseBtn" 
                                    type="button" 
                                    class="close-btn" 
                                    @click="emit('close')"
                                    aria-label="Close"
                                >
                                    <XMarkIcon class="close-icon" />
                                </button>
                            </header>
                        </slot>
                        <slot name="body">
                            <div class="modal-body">
                                <slot name="content"></slot>
                            </div>
                        </slot>
                        <slot name="footer">
                            <div class="modal-footer">
                                <slot name="footerContent"></slot>
                            </div>
                        </slot>
                    </div>
                </div>
            </div>
        </Transition>
    </Teleport>
</template>

<style scoped>
.modal-overlay {
    position: fixed;
    z-index: 2000;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 1rem;
}

.modal-backdrop {
    position: absolute;
    inset: 0;
    background: rgba(0, 0, 0, 0.6);
    backdrop-filter: blur(4px);
}

.modal-container {
    position: relative;
    z-index: 1;
    width: 100%;
    max-width: 420px;
}

.modal-content {
    background: rgba(30, 30, 35, 0.98);
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 1.25rem;
    box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.5);
    overflow: hidden;
}

.modal-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 1.25rem 1.5rem;
    border-bottom: 1px solid rgba(255, 255, 255, 0.08);
}

.modal-title {
    font-family: monospace;
    font-size: 1.25rem;
    font-weight: 600;
    color: #ffffff;
    margin: 0;
}

.close-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 2rem;
    height: 2rem;
    background: rgba(255, 255, 255, 0.05);
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 0.5rem;
    cursor: pointer;
    transition: all 0.2s ease;
}

.close-btn:hover {
    background: rgba(255, 255, 255, 0.1);
}

.close-icon {
    width: 1.25rem;
    height: 1.25rem;
    color: #9ca3af;
}

.modal-body {
    padding: 1.5rem;
}

.modal-footer {
    display: flex;
    gap: 0.75rem;
    padding: 1rem 1.5rem 1.5rem;
}

/* Transition */
.modal-enter-active,
.modal-leave-active {
    transition: opacity 0.2s ease;
}

.modal-enter-active .modal-container,
.modal-leave-active .modal-container {
    transition: transform 0.2s ease, opacity 0.2s ease;
}

.modal-enter-from,
.modal-leave-to {
    opacity: 0;
}

.modal-enter-from .modal-container,
.modal-leave-to .modal-container {
    transform: scale(0.95);
    opacity: 0;
}

@media (max-width: 480px) {
    .modal-container {
        max-width: none;
    }
    
    .modal-header {
        padding: 1rem 1.25rem;
    }
    
    .modal-body {
        padding: 1.25rem;
    }
    
    .modal-footer {
        padding: 0.75rem 1.25rem 1.25rem;
        flex-direction: column;
    }
}
</style>