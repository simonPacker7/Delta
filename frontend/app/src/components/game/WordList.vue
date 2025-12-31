<script setup lang="ts">
import { ref, onUnmounted } from 'vue'
import type { GameTurn } from '@/stores/game'
import wordItem from "./WordItem.vue"

const props = defineProps({
    gameMoves: { type: Array<GameTurn> },
    isGameActive: { type: Boolean },
})

const emit = defineEmits(['timeExpired'])

const secondsLeft = ref(0);
let secondsLeftInterval: number | null = null;

const clearSecondsInterval = () => {
    if (secondsLeftInterval != null) {
        clearInterval(secondsLeftInterval)
    }
}

const restartSecondsInterval = () => {
    clearSecondsInterval()

    secondsLeft.value = 99
    secondsLeftInterval = setInterval(() => {
        if (secondsLeft.value > 0) {
            secondsLeft.value--;
        } else {
            clearSecondsInterval()
            emit("timeExpired")
        }
    }, 1000);
}

const scrollListBottom = (el: any) => {
    el.scrollIntoView({ behavior: 'smooth' })
}

const onNewItem = (el: any) => {
    if (props.isGameActive) {
        restartSecondsInterval()
    } else {
        clearSecondsInterval()
    }
    scrollListBottom(el)
}

const isNewWord = (index: number) => {
    return props.gameMoves != null
        && index == props.gameMoves.length - 1
        && props.isGameActive
}

onUnmounted(() => {
    clearSecondsInterval()
})
</script>

<template>
    <div class="word-container">
        <transition-group name="fade" tag="ul" @enter="onNewItem" class="list-group list-group-flush">
            <transition-group tag="li" v-for="{ Player, Word }, index in gameMoves" :key="index"
                class="list-group-item d-flex align-items-center justify-content-center"
                :class="{ 'new-word-animation': isNewWord(index) }">
                <span class="wordBadge badge text-bg-secondary mono" :key="index">{{ Player.Name }}</span>
                <wordItem :key="index" :word="Word" />
                <span class="secondsRemaining mono" :key="index" v-if="isNewWord(index)">{{ secondsLeft
                }}s</span>
            </transition-group>
        </transition-group>
    </div>
</template>

<style scoped>
.word-container {
    display: flex;
    flex-direction: column-reverse;
    width: 20vw;
    min-width: 480px;
    height: 50vh;
    margin-bottom: 50px;
}

.list-group {
    max-height: 50vh;
    overflow: auto;
    scroll-behavior: smooth;
}

.list-group-item {
    position: relative;
}

.wordBadge {
    position: absolute;
    left: 5%;
}

.secondsRemaining {
    position: absolute;
    left: 90%;
}

/* Fade-in effect */
.fade-enter-active,
.fade-leave-active {
    transition: opacity 1s ease;
}

.fade-enter-from,
.fade-leave-to {
    opacity: 0;
}

/* Hide scrollbar for Chrome, Safari and Opera */
.list-group::-webkit-scrollbar {
    display: none;
}

/* Hide scrollbar for IE, Edge and Firefox */
.list-group {
    -ms-overflow-style: none;
    /* IE and Edge */
    scrollbar-width: none;
    /* Firefox */
}

/* Different color for the last item */
.new-word-animation::before {
    content: "";
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    background-color: #ae9292;
    transform-origin: right;
    z-index: 0;
    animation: colorFade 100s linear, scaleXAnimation 100s linear;
}

@keyframes colorFade {
    0% {
        background-color: #ae9292;
    }

    100% {
        background-color: #953636;
    }
}

@keyframes scaleXAnimation {
    0% {
        transform: scaleX(1);
    }

    100% {
        transform: scaleX(0);
    }
}
</style>