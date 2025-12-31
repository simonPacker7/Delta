<script setup lang="ts">
import { computed } from 'vue';
import Modal from '../modal/Modal.vue';
const props = defineProps({
    winningReason: { type: String },
    isWinningPlayer: { type: Boolean },
    showModal: { type: Boolean, required: true },
})
const emit = defineEmits(['exitGame'])

const onLeaveGame = () => {
    emit('exitGame')
}

const reasonMessage = computed(() => {
    switch (props.winningReason) {
        case "no possible words":
            return "No more valid words";
        case "timelimit":
            return "Time has run out"
        default:
            return "Game is over";
    }
})

const userMessage = computed(() => {
    if (props.isWinningPlayer) {
        return "Congratulations, you won!"
    } else {
        return `You lost. Better luck next time`
    }
})
</script>

<template>
    <div>
        <Modal id="gameOverModal" title="Game Over" :active="showModal" :showCloseBtn="false">
            <template #content>
                <div class="textbox">
                    <span class="reasonMessage">
                        {{ reasonMessage }}
                    </span>
                </div>
                <div class="textbox">
                    <span class="userMessage">
                        {{ userMessage }}
                    </span>
                </div>
            </template>
            <template #footerContent>
                <button class="btn btn-outline-primary" :onclick="onLeaveGame">Leave game</button>
            </template>
        </Modal>
    </div>
</template>

<style scoped>
.textbox {
    display: flex;
    justify-content: center;
    align-items: center;
}

.reasonMessage,
.userMessage {
    font-family: monospace;
    font-weight: 500;
    color: beige;
}
</style>