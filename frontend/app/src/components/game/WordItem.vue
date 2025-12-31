<script setup lang="ts">
const props = defineProps({
    word: { type: String, required: true },
})
</script>

<template>
    <div>
        <span v-for="letter, index in word.split('')" class="wordText mono">
            {{ letter }}
        </span>
    </div>
</template>

<style scoped>
.wordText {
    font-size: 2rem;
    font-weight: 400;
    text-align: center;
    width: 100%;
}

.top-hightlight {
    border-top: 1px solid green;
    border-right: 1px solid green;
    border-left: 1px solid green;
    border-top-left-radius: 1rem;
    border-top-right-radius: 1rem;
}

.full-hightlight {
    border-right: 1px solid green;
    border-left: 1px solid green;
}

.bottom-hightlight {
    border-bottom: 1px solid green;
    border-right: 1px solid green;
    border-left: 1px solid green;
    border-bottom-left-radius: 1rem;
    border-bottom-right-radius: 1rem;
}
</style>