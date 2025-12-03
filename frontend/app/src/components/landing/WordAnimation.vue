<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue';

const words = ['zone', 'gone', 'cone', 'tone'];
const wordIndex = ref(0);
const cursorPosition = ref(0); // we start the cursor at the beginning
const word = ref(words[wordIndex.value].split(''));

const wordMap = {
    zone: 'zonk',
    zonk: 'wonk',
    wonk: 'wink',
    wink: 'pink',
    pink: 'pine',
    pine: 'sine',
    sine: 'vine',
    vine: 'vibe',
    vibe: 'vice',
    vice: 'mice',
    mice: 'mine',
    mine: 'zine',
    zine: 'zone'
}

const nextWord = () => {
    wordIndex.value = (wordIndex.value + 1) % words.length;
    word.value = words[wordIndex.value].split('');
};

const cycleCursor = () => {
    cursorPosition.value = (cursorPosition.value + 1) % 4;
    if (cursorPosition.value === 0) {
        nextWord();
    }
};

let intervalId: number;

onMounted(() => {
    intervalId = setInterval(cycleCursor, 2000); // adjust the speed here
});

onBeforeUnmount(() => {
    clearInterval(intervalId);
})
</script>

<template>
    <div class="word-animation">
        <span v-for="(letter, index) in word" :key="index">
            <span v-if="index === cursorPosition" class="cursor">
                {{ letter }}
            </span>
            <span v-else>
                {{ letter }}
            </span>
        </span>
    </div>
</template>

<style scoped>
.word-animation {
    font-family: monospace;
    font-size: 3rem;
}

.cursor {
    border-bottom: 2px solid white;
}
</style>