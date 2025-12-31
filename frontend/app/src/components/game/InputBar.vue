<script setup lang="ts">
import { ref, computed } from 'vue'
//import { useTooltip } from "@/components/core/tooltip";
const props = defineProps({
    isTurn: { type: Boolean, required: true },
    submitWord: {
        type: Function,
        required: true,
    },
})

const curText = ref('')
const inputRef = ref<Element | null>(null)
//const inputToolTip = useTooltip(inputRef);

const curTurnText = computed(() => props.isTurn ? "Your Turn" : "Opponent Turn")

const SendWord = async () => {
    const error = await props.submitWord(curText.value)
    if (error) {
        //inputToolTip.showToolTip(error)
    } else {
        curText.value = ""
    }
}

</script>

<template>
    <div class="input-container">
        <span class="cur-turn mono">{{ curTurnText }}</span>
        <input ref="inputRef" v-model="curText" class="word-input mono" type="text" :disabled="!isTurn" />
        <button :onclick="SendWord" class="send-word btn mono"
            :class="(isTurn) ? 'btn-outline-success' : 'btn-outline-secondary'" :disabled="!isTurn">Submit</button>
    </div>
</template>

<style scoped>
.input-container {
    display: flex;
    align-items: center;
    justify-content: center;
}

.cur-turn {
    position: absolute;
    left: -175%;
}

.word-input {
    text-align: center;
    font-size: 2rem;
    font-weight: 400;
    width: 5vw;
    min-width: 120px;
    border-width: 0;
    border-bottom-width: 2px;
}

.word-input:focus {
    outline: none;
}

.send-word {
    position: absolute;
    left: 175%;
}
</style>