<script setup lang="ts">
import { ref, watch } from 'vue';
import Modal from '@/components/modal/Modal.vue';
//import { useTooltip } from "@/components/core/tooltip";
const props = defineProps({
    showModal: { type: Boolean, required: true },
    joinGameFunc: { type: Function, required: true },
})
const emit = defineEmits(['cancel', 'joinedGame'])

const joinCode = ref('');
const inputRef = ref<Element | null>(null);
//const inputToolTip = useTooltip(inputRef);

watch(() => props.showModal, () => {
    joinCode.value = '';
//    inputToolTip.hideToolTip()
})

const join = async () => {
    const error = await props.joinGameFunc(joinCode.value)
    if (error != null) {
        //inputToolTip.showToolTip(error)
        console.log(error)
        return
    }
    emit('joinedGame')
}
</script>

<template>
    <div>
        <Modal id="joinGameModal" title="Private Game" :active="showModal" :showCloseBtn="false">
            <template #content>
                <div class="textbox">
                    <label for="joinCode" class="form-label mono">Game Code</label>
                    <input ref="inputRef" v-model="joinCode" id="joinCode" />
                </div>
            </template>
            <template #footerContent>
                <button class="btn btn-outline-primary" :onclick="() => $emit('cancel')">Cancel</button>
                <button class="btn btn-outline-primary" :onclick="join">Join</button>
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