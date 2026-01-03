<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
    word: string
    highlight?: boolean
    previousWord?: string
}>()

// Find which letter changed from the previous word
const changedLetterIndex = computed(() => {
    if (!props.previousWord || props.previousWord.length !== props.word.length) {
        return -1
    }
    for (let i = 0; i < props.word.length; i++) {
        if (props.word[i] !== props.previousWord[i]) {
            return i
        }
    }
    return -1
})

const letters = computed(() => props.word.toUpperCase().split(''))
</script>

<template>
    <div class="word-item">
        <span 
            v-for="(letter, index) in letters" 
            :key="index"
            class="letter-tile"
            :class="{ 
                'changed': changedLetterIndex === index,
                'highlight': highlight
            }"
        >
            {{ letter }}
        </span>
    </div>
</template>

<style scoped>
.word-item {
    display: flex;
    gap: 0.375rem;
    justify-content: center;
}

.letter-tile {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 2.75rem;
    height: 2.75rem;
    font-family: monospace;
    font-size: 1.5rem;
    font-weight: 700;
    color: #e5e7eb;
    background: rgba(255, 255, 255, 0.05);
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 0.5rem;
    transition: all 0.3s ease;
}

.letter-tile.highlight {
    animation: pop-in 0.4s ease;
}

.letter-tile.changed {
    background: rgba(34, 197, 94, 0.15);
    border-color: rgba(34, 197, 94, 0.4);
    color: #22c55e;
}

@keyframes pop-in {
    0% {
        transform: scale(0.8);
        opacity: 0;
    }
    50% {
        transform: scale(1.05);
    }
    100% {
        transform: scale(1);
        opacity: 1;
    }
}

/* Responsive */
@media (max-width: 480px) {
    .word-item {
        gap: 0.25rem;
    }

    .letter-tile {
        width: 2.25rem;
        height: 2.25rem;
        font-size: 1.25rem;
        border-radius: 0.375rem;
    }
}

@media (max-width: 360px) {
    .letter-tile {
        width: 2rem;
        height: 2rem;
        font-size: 1.1rem;
    }
}
</style>