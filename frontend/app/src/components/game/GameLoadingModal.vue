<script setup lang="ts">
import { computed, onUnmounted, ref, watch } from "vue";
import Modal from '@/components/modal/Modal.vue';
const props = defineProps({
    showModal: { type: Boolean, required: true },
    joinCode: { type: String, required: true },
})

const dots = ref('');
let intervalId = -1;

const animateDots = () => {
    let count = 0;
    intervalId = setInterval(() => {
        count = (count + 1) % 4;
        dots.value = '.'.repeat(count);
    }, 500);
}

const stopAnimateDots = () => {
    if (intervalId > -1) {
        clearInterval(intervalId)
        dots.value = ''
        intervalId = -1
    }
}

watch(() => props.showModal, (newV: Boolean) => {
    if (newV) {
        stopAnimateDots()
        animateDots()
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
    return isPrivateGame.value ? "Private Game" : "Matchmaking"
})

</script>

<template>
    <div>
        <Modal id="gameLoadingModal" :title="modalTitle" :active="showModal" :showCloseBtn="false">
            <template #content>
                <div class="textbox">
                    <span v-if="isPrivateGame">
                        JoinCode: {{ joinCode }}
                    </span>
                    <span>
                        Waiting for a player to join<span class="dots">{{ dots }}</span>
                    </span>
                </div>
            </template>
            <template #footerContent>
                <button class="btn btn-outline-primary">Cancel</button>
            </template>
        </Modal>
    </div>
</template>

<style scoped>
.textbox {
    display: flex;
    justify-content: center;
    align-items: center;
    font-family: monospace;
    font-weight: 500;
    color: beige;
    flex-direction: column;
}
</style>