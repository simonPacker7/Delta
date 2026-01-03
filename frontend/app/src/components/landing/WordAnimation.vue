<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue'

const wordChain = [
    'zone', 'zonk', 'wonk', 'wink', 'pink', 'pine', 'sine', 'vine', 'vibe', 'vice', 'mice', 'mine'
]

const wordIndex = ref(0)
const currentWord = ref(wordChain[0].split(''))
const changingIndex = ref(-1)
const isAnimating = ref(false)

const getChangingLetterIndex = (current: string, next: string): number => {
    for (let i = 0; i < current.length; i++) {
        if (current[i] !== next[i]) return i
    }
    return -1
}

const animateToNextWord = async () => {
    const nextIndex = (wordIndex.value + 1) % wordChain.length
    const currentWordStr = wordChain[wordIndex.value]
    const nextWordStr = wordChain[nextIndex]
    
    const letterIndex = getChangingLetterIndex(currentWordStr, nextWordStr)
    
    // Highlight the changing letter
    isAnimating.value = true
    changingIndex.value = letterIndex
    
    // Wait, then change the letter
    await new Promise(resolve => setTimeout(resolve, 600))
    
    currentWord.value = nextWordStr.split('')
    wordIndex.value = nextIndex
    
    // Brief pause showing the new letter highlighted
    await new Promise(resolve => setTimeout(resolve, 400))
    
    changingIndex.value = -1
    isAnimating.value = false
}

let intervalId: number

onMounted(() => {
    intervalId = setInterval(animateToNextWord, 2000)
})

onBeforeUnmount(() => {
    clearInterval(intervalId)
})
</script>

<template>
    <div class="word-animation">
        <span
            v-for="(letter, index) in currentWord"
            :key="index"
            :class="['letter', { 
                'changing': changingIndex === index,
                'highlight': changingIndex === index && isAnimating
            }]"
        >
            {{ letter }}
        </span>
    </div>
</template>

<style scoped>
.word-animation {
    display: flex;
    gap: 0.25rem;
    font-family: monospace;
    font-size: clamp(2.5rem, 6vw, 4rem);
    font-weight: 600;
    letter-spacing: 0.1em;
}

.letter {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 1.5em;
    height: 1.5em;
    border-radius: 0.375rem;
    background: rgba(255, 255, 255, 0.05);
    border: 1px solid rgba(255, 255, 255, 0.1);
    transition: all 0.3s ease;
    text-transform: uppercase;
}

.letter.changing {
    border-color: #6366f1;
    background: rgba(99, 102, 241, 0.15);
}

.letter.highlight {
    animation: pulse 0.6s ease-in-out;
    box-shadow: 0 0 20px rgba(99, 102, 241, 0.5);
}

@keyframes pulse {
    0%, 100% {
        transform: scale(1);
    }
    50% {
        transform: scale(1.1);
        background: rgba(99, 102, 241, 0.3);
    }
}

@media (max-width: 640px) {
    .word-animation {
        gap: 0.15rem;
    }

    .letter {
        width: 1.4em;
        height: 1.4em;
        border-radius: 0.25rem;
    }
}
</style>