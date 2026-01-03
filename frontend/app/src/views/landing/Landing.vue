<script setup lang="ts">
import { ref } from 'vue'
import WordAnimation from '@/components/landing/WordAnimation.vue'
import RegisterView from '@/components/landing/RegisterView.vue'
import LoginView from '@/components/landing/LoginView.vue'
import { ArrowRightIcon, SparklesIcon, UsersIcon, ClockIcon } from '@heroicons/vue/24/outline'

const activeTab = ref<'login' | 'register'>('login')
</script>

<template>
    <div class="landing">
        <!-- Hero Section -->
        <section class="hero">
            <div class="hero-content">
                <div class="hero-text">
                    <div class="logo-badge">
                        <SparklesIcon class="badge-icon" />
                        <span>Word Game</span>
                    </div>
                    <h1 class="hero-title">
                        <span class="title-delta">Delta</span>
                    </h1>
                    <p class="hero-tagline">One letter at a time</p>
                    <div class="word-animation-wrapper">
                        <WordAnimation />
                    </div>
                    <p class="hero-description">
                        Challenge your vocabulary in this word transformation game
                    </p>

                    <!-- Feature highlights -->
                    <div class="features">
                        <div class="feature">
                            <UsersIcon class="feature-icon" />
                            <span>Real-time multiplayer</span>
                        </div>
                        <div class="feature">
                            <ClockIcon class="feature-icon" />
                            <span>Fast-paced turns</span>
                        </div>
                    </div>
                </div>

                <!-- Auth Card -->
                <div class="auth-section">
                    <div class="auth-card">
                        <div class="auth-tabs">
                            <button 
                                :class="['tab', { active: activeTab === 'login' }]"
                                @click="activeTab = 'login'"
                            >
                                Sign In
                            </button>
                            <button 
                                :class="['tab', { active: activeTab === 'register' }]"
                                @click="activeTab = 'register'"
                            >
                                Create Account
                            </button>
                        </div>
                        <div class="auth-content">
                            <Transition name="fade" mode="out-in">
                                <LoginView v-if="activeTab === 'login'" key="login" />
                                <RegisterView v-else key="register" />
                            </Transition>
                        </div>
                        <div class="auth-footer">
                            <p v-if="activeTab === 'login'">
                                New to Delta?
                                <button class="link-btn" @click="activeTab = 'register'">
                                    Create an account
                                    <ArrowRightIcon class="link-icon" />
                                </button>
                            </p>
                            <p v-else>
                                Already have an account?
                                <button class="link-btn" @click="activeTab = 'login'">
                                    Sign in
                                    <ArrowRightIcon class="link-icon" />
                                </button>
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </section>

        <!-- Background decorations -->
        <div class="bg-decoration">
            <div class="gradient-orb orb-1"></div>
            <div class="gradient-orb orb-2"></div>
            <div class="gradient-orb orb-3"></div>
        </div>
    </div>
</template>

<style scoped>
.landing {
    min-height: 100vh;
    width: 100%;
    position: relative;
    overflow: hidden;
}

/* Hero Section */
.hero {
    position: relative;
    z-index: 10;
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 2rem;
    box-sizing: border-box;
}

.hero-content {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 4rem;
    max-width: 1200px;
    width: 100%;
    align-items: center;
}

/* Left side - Hero text */
.hero-text {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
}

.logo-badge {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    background: rgba(99, 102, 241, 0.15);
    border: 1px solid rgba(99, 102, 241, 0.3);
    padding: 0.5rem 1rem;
    border-radius: 2rem;
    width: fit-content;
    font-size: 0.875rem;
    color: #a5b4fc;
    font-family: monospace;
}

.badge-icon {
    width: 1rem;
    height: 1rem;
}

.hero-title {
    font-size: clamp(3rem, 8vw, 5rem);
    font-weight: 700;
    line-height: 1.1;
    margin: 0;
}

.title-delta {
    background: linear-gradient(135deg, #ffffff 0%, #a5b4fc 50%, #6366f1 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    background-clip: text;
    font-family: monospace;
}

.hero-tagline {
    font-size: 1.5rem;
    color: #9ca3af;
    font-family: monospace;
    margin: 0;
}

.word-animation-wrapper {
    background: rgba(255, 255, 255, 0.05);
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 1rem;
    padding: 1.5rem 2rem;
    width: fit-content;
}

.hero-description {
    font-size: 1.125rem;
    color: #9ca3af;
    line-height: 1.7;
    max-width: 500px;
    margin: 0;
}

/* Features */
.features {
    display: flex;
    gap: 2rem;
    margin-top: 1rem;
}

.feature {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    color: #d1d5db;
    font-size: 0.9rem;
}

.feature-icon {
    width: 1.25rem;
    height: 1.25rem;
    color: #6366f1;
}

/* Auth Section */
.auth-section {
    display: flex;
    justify-content: center;
}

.auth-card {
    background: rgba(255, 255, 255, 0.03);
    backdrop-filter: blur(20px);
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 1.5rem;
    padding: 2rem;
    width: 100%;
    max-width: 420px;
    box-shadow: 
        0 25px 50px -12px rgba(0, 0, 0, 0.5),
        0 0 0 1px rgba(255, 255, 255, 0.05) inset;
}

.auth-tabs {
    display: flex;
    background: rgba(255, 255, 255, 0.05);
    border-radius: 0.75rem;
    padding: 0.25rem;
    margin-bottom: 1.5rem;
}

.tab {
    flex: 1;
    padding: 0.75rem 1rem;
    border: none;
    background: transparent;
    color: #9ca3af;
    font-family: monospace;
    font-size: 0.9rem;
    border-radius: 0.5rem;
    cursor: pointer;
    transition: all 0.2s ease;
}

.tab:hover {
    color: #ffffff;
}

.tab.active {
    background: rgba(99, 102, 241, 0.2);
    color: #ffffff;
    box-shadow: 0 2px 8px rgba(99, 102, 241, 0.3);
}

.auth-content {
    min-height: 280px;
}

.auth-footer {
    margin-top: 1.5rem;
    padding-top: 1.5rem;
    border-top: 1px solid rgba(255, 255, 255, 0.1);
    text-align: center;
}

.auth-footer p {
    color: #9ca3af;
    font-size: 0.9rem;
    margin: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
    flex-wrap: wrap;
}

.link-btn {
    background: none;
    border: none;
    color: #a5b4fc;
    font-family: monospace;
    font-size: 0.9rem;
    cursor: pointer;
    display: inline-flex;
    align-items: center;
    gap: 0.25rem;
    transition: color 0.2s ease;
}

.link-btn:hover {
    color: #6366f1;
}

.link-icon {
    width: 0.875rem;
    height: 0.875rem;
}

/* Transition animations */
.fade-enter-active,
.fade-leave-active {
    transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
    opacity: 0;
}

/* Background decorations */
.bg-decoration {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    pointer-events: none;
    z-index: 1;
    overflow: hidden;
}

.gradient-orb {
    position: absolute;
    border-radius: 50%;
    filter: blur(80px);
    opacity: 0.4;
}

.orb-1 {
    width: 600px;
    height: 600px;
    background: radial-gradient(circle, #6366f1 0%, transparent 70%);
    top: -200px;
    right: -100px;
    animation: float 20s ease-in-out infinite;
}

.orb-2 {
    width: 400px;
    height: 400px;
    background: radial-gradient(circle, #8b5cf6 0%, transparent 70%);
    bottom: -100px;
    left: -100px;
    animation: float 15s ease-in-out infinite reverse;
}

.orb-3 {
    width: 300px;
    height: 300px;
    background: radial-gradient(circle, #06b6d4 0%, transparent 70%);
    top: 50%;
    left: 30%;
    animation: float 18s ease-in-out infinite;
    opacity: 0.2;
}

@keyframes float {
    0%, 100% {
        transform: translate(0, 0) scale(1);
    }
    33% {
        transform: translate(30px, -30px) scale(1.05);
    }
    66% {
        transform: translate(-20px, 20px) scale(0.95);
    }
}

/* Responsive Design */
@media (max-width: 1024px) {
    .hero-content {
        grid-template-columns: 1fr;
        gap: 3rem;
        text-align: center;
    }

    .hero-text {
        align-items: center;
    }

    .hero-description {
        max-width: 600px;
    }

    .features {
        justify-content: center;
    }

    .word-animation-wrapper {
        margin: 0 auto;
    }
}

@media (max-width: 640px) {
    .hero {
        padding: 1rem;
    }

    .hero-content {
        gap: 2rem;
    }

    .hero-tagline {
        font-size: 1.25rem;
    }

    .hero-description {
        font-size: 1rem;
    }

    .features {
        flex-direction: column;
        gap: 1rem;
    }

    .auth-card {
        padding: 1.5rem;
        border-radius: 1rem;
    }

    .auth-tabs {
        flex-direction: column;
    }

    .tab {
        padding: 0.875rem;
    }

    .word-animation-wrapper {
        padding: 1rem 1.5rem;
    }

    .orb-1 {
        width: 300px;
        height: 300px;
    }

    .orb-2 {
        width: 200px;
        height: 200px;
    }

    .orb-3 {
        display: none;
    }
}

@media (max-width: 400px) {
    .auth-card {
        padding: 1rem;
    }

    .logo-badge {
        font-size: 0.75rem;
        padding: 0.375rem 0.75rem;
    }
}
</style>