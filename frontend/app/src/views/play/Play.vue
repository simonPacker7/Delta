<script setup lang="ts">
import { ref } from 'vue';
import router from '@/router';
import { useGameStore } from '@/stores/game'

const gameStore = useGameStore()

const showJoinPrivateGameModal = ref(false)

const findGame = async () => {
    await gameStore.FindGame()
    await router.push("/play/online")
}

const createPrivateGame = async () => {
    // await gameStore.CreatePrivateGame()
    // await router.push("/play/online")
}

const joinPrivateGame = async () => {
    showJoinPrivateGameModal.value = false
    await router.push("/play/online")
}

const localGame = async () => {
    await router.push("/play/local")
}

</script>

<template>
    <div class="play-container">
        <div>
            <h3>Online</h3>
            <h6>Play against others</h6>
            <div class="play-btn-container">
                <button class="btn btn-outline-primary" :onclick="findGame">
                    Find a match
                </button>
            </div>
        </div>
        <div>
            <h3>Private</h3>
            <h6>Challenge your friends</h6>
            <div class="play-btn-container">
                <button class="btn btn-outline-primary" :onclick="createPrivateGame">
                    Create a game
                </button>
                <button class="btn btn-outline-primary" :onclick="() => showJoinPrivateGameModal = true">
                    Join with code
                </button>
            </div>
        </div>
        <div>
            <h3>Single Player</h3>
            <h6>Practice your skills</h6>
            <div class="play-btn-container">
                <button class="btn btn-outline-primary" :onclick="localGame">
                    Against CPU
                </button>
            </div>
        </div>
    </div>
</template>

<style scoped>
.play-container {
    display: flex;
    flex-direction: row;
    justify-content: space-around;
}

.play-container>div {
    display: flex;
    justify-content: flex-start;
    flex-direction: column;
    align-items: center;
    width: 30vw;
    border-right: 1px dashed;
    font-family: monospace;
}

.play-container>div:last-child {
    border-right: 0px solid;
}

.play-btn-container {
    display: flex;
    justify-content: flex-start;
    flex-direction: column;
}

.play-btn-container>button {
    height: min-content;
    font-family: monospace;
    margin-top: 15px;
}
</style>