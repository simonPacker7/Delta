<script setup lang="ts">
import { ref } from 'vue'
import router from '@/router'
import { useGameStore } from '@/stores/game'
import joinWithCodeModal from "@/components/game/JoinWithCodeModal.vue"
import { 
    GlobeAltIcon, 
    UserGroupIcon, 
    CpuChipIcon,
    MagnifyingGlassIcon,
    PlusIcon,
    KeyIcon,
    PlayIcon,
    ArrowRightIcon
} from '@heroicons/vue/24/outline'

const gameStore = useGameStore()

const showJoinPrivateGameModal = ref(false)
const isLoadingQuickMatch = ref(false)
const isLoadingCreatePrivate = ref(false)

const findGame = async () => {
    isLoadingQuickMatch.value = true
    try {
        await gameStore.FindGame()
        await router.push("/play/online")
    } finally {
        isLoadingQuickMatch.value = false
    }
}

const createPrivateGame = async () => {
    isLoadingCreatePrivate.value = true
    try {
        await gameStore.CreatePrivateGame()
        await router.push("/play/online")
    } finally {
        isLoadingCreatePrivate.value = false
    }
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
    <div class="play-page">
        <!-- Page Header -->
        <div class="page-header">
            <h1 class="page-title">Choose Your Game Mode</h1>
            <p class="page-subtitle">Select how you want to play Delta</p>
        </div>

        <!-- Game Mode Cards -->
        <div class="game-modes">
            <!-- Online Matchmaking -->
            <div class="mode-card featured">
                <div class="mode-header">
                    <div class="mode-icon-wrapper online">
                        <GlobeAltIcon class="mode-icon" />
                    </div>
                    <div class="mode-badge">Most Popular</div>
                </div>
                <h2 class="mode-title">Online Match</h2>
                <p class="mode-description">
                    Get matched with a random opponent and compete in real-time. 
                    Test your skills against players from around the world.
                </p>
                <div class="mode-features">
                    <span class="feature-tag">
                        <PlayIcon class="feature-icon" />
                        Quick matchmaking
                    </span>
                    <span class="feature-tag">
                        <ArrowRightIcon class="feature-icon" />
                        Ranked play
                    </span>
                </div>
                <button 
                    class="mode-btn primary" 
                    @click="findGame"
                    :disabled="isLoadingQuickMatch"
                >
                    <MagnifyingGlassIcon v-if="!isLoadingQuickMatch" class="btn-icon" />
                    <span v-if="isLoadingQuickMatch" class="loading-spinner"></span>
                    <span>{{ isLoadingQuickMatch ? 'Searching...' : 'Find a Match' }}</span>
                </button>
            </div>

            <!-- Private Game -->
            <div class="mode-card">
                <div class="mode-header">
                    <div class="mode-icon-wrapper private">
                        <UserGroupIcon class="mode-icon" />
                    </div>
                </div>
                <h2 class="mode-title">Private Game</h2>
                <p class="mode-description">
                    Create a private room and invite your friends with a code. 
                    Perfect for friendly matches and practice sessions.
                </p>
                <div class="mode-features">
                    <span class="feature-tag">
                        <KeyIcon class="feature-icon" />
                        Share game code
                    </span>
                </div>
                <div class="mode-actions">
                    <button 
                        class="mode-btn secondary" 
                        @click="createPrivateGame"
                        :disabled="isLoadingCreatePrivate"
                    >
                        <PlusIcon v-if="!isLoadingCreatePrivate" class="btn-icon" />
                        <span v-if="isLoadingCreatePrivate" class="loading-spinner"></span>
                        <span>{{ isLoadingCreatePrivate ? 'Creating...' : 'Create Game' }}</span>
                    </button>
                    <button 
                        class="mode-btn outline" 
                        @click="() => showJoinPrivateGameModal = true"
                    >
                        <KeyIcon class="btn-icon" />
                        <span>Join with Code</span>
                    </button>
                </div>
            </div>

            <!-- Single Player -->
            <div class="mode-card">
                <div class="mode-header">
                    <div class="mode-icon-wrapper single">
                        <CpuChipIcon class="mode-icon" />
                    </div>
                    <div class="mode-badge coming-soon">Coming Soon</div>
                </div>
                <h2 class="mode-title">Single Player</h2>
                <p class="mode-description">
                    Practice against the CPU to sharpen your word skills. 
                    Multiple difficulty levels to challenge yourself.
                </p>
                <div class="mode-features">
                    <span class="feature-tag">
                        <CpuChipIcon class="feature-icon" />
                        AI opponent
                    </span>
                </div>
                <button class="mode-btn disabled" disabled @click="localGame">
                    <PlayIcon class="btn-icon" />
                    <span>Coming Soon</span>
                </button>
            </div>
        </div>

        <!-- Join with Code Modal -->
        <joinWithCodeModal 
            :show-modal="showJoinPrivateGameModal"
            :join-game-func="(code: string) => gameStore.JoinPrivateGame(code)"
            @cancel="() => showJoinPrivateGameModal = false" 
            @joined-game="joinPrivateGame" 
        />

        <!-- Background decoration -->
        <div class="bg-decoration">
            <div class="gradient-orb orb-1"></div>
            <div class="gradient-orb orb-2"></div>
        </div>
    </div>
</template>

<style scoped>
.play-page {
    position: relative;
    min-height: calc(100vh - 70px);
    padding: 2rem;
    max-width: 1200px;
    margin: 0 auto;
}

/* Page Header */
.page-header {
    position: relative;
    z-index: 10;
    text-align: center;
    padding: 2rem 0 3rem;
}

.page-title {
    font-size: clamp(1.75rem, 4vw, 2.5rem);
    font-weight: 700;
    color: #ffffff;
    margin: 0 0 0.75rem;
    font-family: monospace;
}

.page-subtitle {
    font-size: 1.125rem;
    color: #9ca3af;
    margin: 0;
}

/* Game Modes Grid */
.game-modes {
    position: relative;
    z-index: 10;
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 1.5rem;
}

/* Mode Card */
.mode-card {
    display: flex;
    flex-direction: column;
    background: rgba(255, 255, 255, 0.03);
    border: 1px solid rgba(255, 255, 255, 0.08);
    border-radius: 1.25rem;
    padding: 1.75rem;
    transition: all 0.3s ease;
}

.mode-card:hover {
    background: rgba(255, 255, 255, 0.05);
    border-color: rgba(255, 255, 255, 0.12);
    transform: translateY(-4px);
}

.mode-card.featured {
    background: linear-gradient(135deg, rgba(99, 102, 241, 0.1) 0%, rgba(139, 92, 246, 0.1) 100%);
    border-color: rgba(99, 102, 241, 0.25);
}

.mode-card.featured:hover {
    border-color: rgba(99, 102, 241, 0.4);
}

.mode-header {
    display: flex;
    align-items: flex-start;
    justify-content: space-between;
    margin-bottom: 1.25rem;
}

.mode-icon-wrapper {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 3.5rem;
    height: 3.5rem;
    border-radius: 1rem;
}

.mode-icon-wrapper.online {
    background: rgba(99, 102, 241, 0.15);
}

.mode-icon-wrapper.private {
    background: rgba(34, 197, 94, 0.15);
}

.mode-icon-wrapper.single {
    background: rgba(249, 115, 22, 0.15);
}

.mode-icon {
    width: 1.75rem;
    height: 1.75rem;
}

.mode-icon-wrapper.online .mode-icon {
    color: #6366f1;
}

.mode-icon-wrapper.private .mode-icon {
    color: #22c55e;
}

.mode-icon-wrapper.single .mode-icon {
    color: #f97316;
}

.mode-badge {
    font-size: 0.7rem;
    font-weight: 600;
    text-transform: uppercase;
    letter-spacing: 0.05em;
    padding: 0.375rem 0.75rem;
    background: rgba(99, 102, 241, 0.2);
    color: #a5b4fc;
    border-radius: 1rem;
    font-family: monospace;
}

.mode-badge.coming-soon {
    background: rgba(249, 115, 22, 0.15);
    color: #fb923c;
}

.mode-title {
    font-size: 1.35rem;
    font-weight: 600;
    color: #ffffff;
    margin: 0 0 0.75rem;
    font-family: monospace;
}

.mode-description {
    font-size: 0.9rem;
    color: #9ca3af;
    line-height: 1.6;
    margin: 0 0 1.25rem;
    flex: 1;
}

.mode-features {
    display: flex;
    flex-wrap: wrap;
    gap: 0.5rem;
    margin-bottom: 1.5rem;
}

.feature-tag {
    display: inline-flex;
    align-items: center;
    gap: 0.375rem;
    font-size: 0.75rem;
    color: #6b7280;
    background: rgba(255, 255, 255, 0.05);
    padding: 0.375rem 0.625rem;
    border-radius: 0.375rem;
    font-family: monospace;
}

.feature-icon {
    width: 0.875rem;
    height: 0.875rem;
}

.mode-actions {
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
}

/* Mode Buttons */
.mode-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
    padding: 0.875rem 1.25rem;
    border-radius: 0.75rem;
    font-size: 0.95rem;
    font-weight: 600;
    font-family: monospace;
    cursor: pointer;
    transition: all 0.2s ease;
    border: none;
    width: 100%;
}

.mode-btn.primary {
    background: linear-gradient(135deg, #6366f1 0%, #8b5cf6 100%);
    color: #ffffff;
}

.mode-btn.primary:hover:not(:disabled) {
    box-shadow: 0 8px 20px rgba(99, 102, 241, 0.4);
    transform: translateY(-2px);
}

.mode-btn.secondary {
    background: rgba(34, 197, 94, 0.15);
    color: #22c55e;
    border: 1px solid rgba(34, 197, 94, 0.3);
}

.mode-btn.secondary:hover:not(:disabled) {
    background: rgba(34, 197, 94, 0.25);
    border-color: rgba(34, 197, 94, 0.5);
}

.mode-btn.outline {
    background: transparent;
    color: #9ca3af;
    border: 1px solid rgba(255, 255, 255, 0.15);
}

.mode-btn.outline:hover {
    background: rgba(255, 255, 255, 0.05);
    border-color: rgba(255, 255, 255, 0.25);
    color: #ffffff;
}

.mode-btn.disabled {
    background: rgba(255, 255, 255, 0.05);
    color: #4b5563;
    cursor: not-allowed;
}

.mode-btn:disabled {
    opacity: 0.7;
    cursor: not-allowed;
    transform: none !important;
}

.btn-icon {
    width: 1.25rem;
    height: 1.25rem;
}

.loading-spinner {
    width: 1.25rem;
    height: 1.25rem;
    border: 2px solid rgba(255, 255, 255, 0.3);
    border-top-color: #ffffff;
    border-radius: 50%;
    animation: spin 0.8s linear infinite;
}

@keyframes spin {
    to {
        transform: rotate(360deg);
    }
}

/* Background decoration */
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
    filter: blur(100px);
    opacity: 0.25;
}

.orb-1 {
    width: 500px;
    height: 500px;
    background: radial-gradient(circle, #6366f1 0%, transparent 70%);
    top: 0;
    right: -150px;
}

.orb-2 {
    width: 400px;
    height: 400px;
    background: radial-gradient(circle, #22c55e 0%, transparent 70%);
    bottom: -100px;
    left: -100px;
}

/* Responsive Design */
@media (max-width: 1024px) {
    .game-modes {
        grid-template-columns: 1fr;
        max-width: 500px;
        margin: 0 auto;
    }

    .mode-card.featured {
        order: -1;
    }
}

@media (max-width: 640px) {
    .play-page {
        padding: 1.5rem 1rem;
    }

    .page-header {
        padding: 1rem 0 2rem;
    }

    .page-subtitle {
        font-size: 1rem;
    }

    .mode-card {
        padding: 1.5rem;
    }

    .mode-actions {
        flex-direction: column;
    }

    .orb-1, .orb-2 {
        width: 250px;
        height: 250px;
    }
}

@media (max-width: 480px) {
    .play-page {
        min-height: calc(100vh - 60px);
    }

    .mode-icon-wrapper {
        width: 3rem;
        height: 3rem;
    }

    .mode-icon {
        width: 1.5rem;
        height: 1.5rem;
    }

    .mode-title {
        font-size: 1.2rem;
    }
}
</style>