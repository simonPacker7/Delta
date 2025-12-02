<script setup lang="ts">
import WordAnimation from '@/components/landing/WordAnimation.vue';
import RegisterView from '@/components/landing/RegisterView.vue';
import LoginView from '@/components/landing/LoginView.vue';
</script>

<template>
    <div class="landing">
        <div class="head">
            <h2 class="motto text-muted">The mathematical symbol for change</h2>
            <WordAnimation></WordAnimation>
        </div>
        <div class="btnView">
            <div>
                <LoginView></LoginView>
            </div>
            <div>
                <RegisterView></RegisterView>
            </div>
        </div>
    </div>
</template>

<style scoped>
.landing {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;
    height: 90vh;
}

.head {
    display: flex !important;
    align-items: center;
    flex-direction: column;
    margin-bottom: 10rem;
}

.title {
    font-family: "Menlo";
    font-weight: 600;
    font-size: 10rem;
}

.motto {
    font-family: "Menlo";
    font-weight: 300;
    font-size: 3rem;
}

.btnView {
    display: flex;
    justify-content: space-between;
    gap: 2rem;
    padding-bottom: 10rem;
}

.btnView>button {
    min-width: 8rem;
    min-height: 3rem;
    font-family: monospace;
    font-size: 2rem;
}
</style>