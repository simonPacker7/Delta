<script setup lang="ts">
import { useSessionStore } from '@/stores/session'
import { useRouter } from 'vue-router'
import { 
    PlayIcon, 
    UserGroupIcon, 
    TrophyIcon, 
    ChartBarIcon,
    SparklesIcon,
    ArrowRightIcon,
    ClockIcon,
    FireIcon
} from '@heroicons/vue/24/outline'

const sessionStore = useSessionStore()
const router = useRouter()

const navigateTo = (path: string) => {
    router.push(path)
}

// Placeholder stats - these would come from an API in a real implementation
const stats = {
    gamesPlayed: 42,
    wins: 28,
    winRate: 67,
    currentStreak: 5
}
</script>

<template>
    <div class="home">
        <!-- Welcome Section -->
        <section class="welcome-section">
            <div class="welcome-content">
                <div class="welcome-badge">
                    <SparklesIcon class="badge-icon" />
                    <span>Welcome back</span>
                </div>
                <h1 class="welcome-title">
                    Hey, <span class="username">{{ sessionStore.displayName }}</span>!
                </h1>
                <p class="welcome-subtitle">
                    Jump into a game or check your stats
                </p>
            </div>
        </section>

        <!-- Quick Stats Cards -->
        <section class="stats-section">
            <div class="stats-grid">
                <div class="stat-card">
                    <div class="stat-icon-wrapper games">
                        <PlayIcon class="stat-icon" />
                    </div>
                    <div class="stat-info">
                        <span class="stat-value">{{ stats.gamesPlayed }}</span>
                        <span class="stat-label">Games Played</span>
                    </div>
                </div>
                <div class="stat-card">
                    <div class="stat-icon-wrapper wins">
                        <TrophyIcon class="stat-icon" />
                    </div>
                    <div class="stat-info">
                        <span class="stat-value">{{ stats.wins }}</span>
                        <span class="stat-label">Victories</span>
                    </div>
                </div>
                <div class="stat-card">
                    <div class="stat-icon-wrapper rate">
                        <ChartBarIcon class="stat-icon" />
                    </div>
                    <div class="stat-info">
                        <span class="stat-value">{{ stats.winRate }}%</span>
                        <span class="stat-label">Win Rate</span>
                    </div>
                </div>
                <div class="stat-card">
                    <div class="stat-icon-wrapper streak">
                        <FireIcon class="stat-icon" />
                    </div>
                    <div class="stat-info">
                        <span class="stat-value">{{ stats.currentStreak }}</span>
                        <span class="stat-label">Win Streak</span>
                    </div>
                </div>
            </div>
        </section>

        <!-- Action Cards -->
        <section class="actions-section">
            <h2 class="section-title">Let's Play</h2>
            <div class="actions-grid">
                <!-- Quick Match Card -->
                <button class="action-card primary" @click="navigateTo('/play')">
                    <div class="action-icon-wrapper">
                        <PlayIcon class="action-icon" />
                    </div>
                    <div class="action-content">
                        <h3 class="action-title">Quick Match</h3>
                        <p class="action-description">
                            Jump into a game against a random opponent
                        </p>
                    </div>
                    <ArrowRightIcon class="action-arrow" />
                </button>

                <!-- Challenge Friends Card -->
                <button class="action-card" @click="navigateTo('/play')">
                    <div class="action-icon-wrapper friends">
                        <UserGroupIcon class="action-icon" />
                    </div>
                    <div class="action-content">
                        <h3 class="action-title">Challenge Friends</h3>
                        <p class="action-description">
                            Create a private game and invite your friends
                        </p>
                    </div>
                    <ArrowRightIcon class="action-arrow" />
                </button>

                <!-- Leaderboard Card -->
                <button class="action-card" @click="navigateTo('/leaderboard')">
                    <div class="action-icon-wrapper leaderboard">
                        <TrophyIcon class="action-icon" />
                    </div>
                    <div class="action-content">
                        <h3 class="action-title">Leaderboard</h3>
                        <p class="action-description">
                            See how you rank against other players
                        </p>
                    </div>
                    <ArrowRightIcon class="action-arrow" />
                </button>

                <!-- Statistics Card -->
                <button class="action-card" @click="navigateTo('/statistics')">
                    <div class="action-icon-wrapper stats">
                        <ChartBarIcon class="action-icon" />
                    </div>
                    <div class="action-content">
                        <h3 class="action-title">My Statistics</h3>
                        <p class="action-description">
                            View your detailed game history and stats
                        </p>
                    </div>
                    <ArrowRightIcon class="action-arrow" />
                </button>
            </div>
        </section>

        <!-- Recent Activity Section (Placeholder) -->
        <section class="activity-section">
            <h2 class="section-title">Recent Games</h2>
            <div class="activity-placeholder">
                <ClockIcon class="placeholder-icon" />
                <p>Your recent games will appear here</p>
            </div>
        </section>

        <!-- Background decoration -->
        <div class="bg-decoration">
            <div class="gradient-orb orb-1"></div>
            <div class="gradient-orb orb-2"></div>
        </div>
    </div>
</template>

<style scoped>
.home {
    position: relative;
    min-height: calc(100vh - 70px);
    padding: 2rem;
    max-width: 1200px;
    margin: 0 auto;
    overflow: hidden;
}

/* Welcome Section */
.welcome-section {
    position: relative;
    z-index: 10;
    padding: 2rem 0 3rem;
}

.welcome-content {
    display: flex;
    flex-direction: column;
    gap: 1rem;
}

.welcome-badge {
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

.welcome-title {
    font-size: clamp(2rem, 5vw, 3rem);
    font-weight: 700;
    margin: 0;
    color: #ffffff;
    font-family: monospace;
}

.username {
    background: linear-gradient(135deg, #a5b4fc 0%, #6366f1 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    background-clip: text;
}

.welcome-subtitle {
    font-size: 1.125rem;
    color: #9ca3af;
    margin: 0;
    max-width: 500px;
}

/* Stats Section */
.stats-section {
    position: relative;
    z-index: 10;
    margin-bottom: 3rem;
}

.stats-grid {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: 1rem;
}

.stat-card {
    display: flex;
    align-items: center;
    gap: 1rem;
    background: rgba(255, 255, 255, 0.03);
    border: 1px solid rgba(255, 255, 255, 0.08);
    border-radius: 1rem;
    padding: 1.25rem;
    transition: all 0.2s ease;
}

.stat-card:hover {
    background: rgba(255, 255, 255, 0.05);
    border-color: rgba(255, 255, 255, 0.12);
}

.stat-icon-wrapper {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 3rem;
    height: 3rem;
    border-radius: 0.75rem;
    flex-shrink: 0;
}

.stat-icon-wrapper.games {
    background: rgba(99, 102, 241, 0.15);
    color: #6366f1;
}

.stat-icon-wrapper.wins {
    background: rgba(34, 197, 94, 0.15);
    color: #22c55e;
}

.stat-icon-wrapper.rate {
    background: rgba(168, 85, 247, 0.15);
    color: #a855f7;
}

.stat-icon-wrapper.streak {
    background: rgba(249, 115, 22, 0.15);
    color: #f97316;
}

.stat-icon {
    width: 1.5rem;
    height: 1.5rem;
}

.stat-info {
    display: flex;
    flex-direction: column;
    gap: 0.25rem;
}

.stat-value {
    font-size: 1.5rem;
    font-weight: 700;
    color: #ffffff;
    font-family: monospace;
    line-height: 1;
}

.stat-label {
    font-size: 0.8rem;
    color: #6b7280;
    font-family: monospace;
}

/* Actions Section */
.actions-section {
    position: relative;
    z-index: 10;
    margin-bottom: 3rem;
}

.section-title {
    font-size: 1.25rem;
    font-weight: 600;
    color: #e5e7eb;
    margin: 0 0 1.25rem;
    font-family: monospace;
}

.actions-grid {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 1rem;
}

.action-card {
    display: flex;
    align-items: center;
    gap: 1rem;
    background: rgba(255, 255, 255, 0.03);
    border: 1px solid rgba(255, 255, 255, 0.08);
    border-radius: 1rem;
    padding: 1.5rem;
    cursor: pointer;
    transition: all 0.2s ease;
    text-align: left;
    width: 100%;
}

.action-card:hover {
    background: rgba(255, 255, 255, 0.06);
    border-color: rgba(255, 255, 255, 0.15);
    transform: translateY(-2px);
}

.action-card.primary {
    background: linear-gradient(135deg, rgba(99, 102, 241, 0.15) 0%, rgba(139, 92, 246, 0.15) 100%);
    border-color: rgba(99, 102, 241, 0.3);
}

.action-card.primary:hover {
    background: linear-gradient(135deg, rgba(99, 102, 241, 0.25) 0%, rgba(139, 92, 246, 0.25) 100%);
    border-color: rgba(99, 102, 241, 0.5);
}

.action-icon-wrapper {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 3.5rem;
    height: 3.5rem;
    background: rgba(99, 102, 241, 0.15);
    border-radius: 0.875rem;
    flex-shrink: 0;
}

.action-icon-wrapper.friends {
    background: rgba(34, 197, 94, 0.15);
}

.action-icon-wrapper.leaderboard {
    background: rgba(249, 115, 22, 0.15);
}

.action-icon-wrapper.stats {
    background: rgba(168, 85, 247, 0.15);
}

.action-icon {
    width: 1.75rem;
    height: 1.75rem;
    color: #a5b4fc;
}

.action-icon-wrapper.friends .action-icon {
    color: #22c55e;
}

.action-icon-wrapper.leaderboard .action-icon {
    color: #f97316;
}

.action-icon-wrapper.stats .action-icon {
    color: #a855f7;
}

.action-content {
    flex: 1;
    min-width: 0;
}

.action-title {
    font-size: 1.1rem;
    font-weight: 600;
    color: #ffffff;
    margin: 0 0 0.375rem;
    font-family: monospace;
}

.action-description {
    font-size: 0.875rem;
    color: #6b7280;
    margin: 0;
    line-height: 1.4;
}

.action-arrow {
    width: 1.25rem;
    height: 1.25rem;
    color: #4b5563;
    flex-shrink: 0;
    transition: all 0.2s ease;
}

.action-card:hover .action-arrow {
    color: #9ca3af;
    transform: translateX(4px);
}

/* Activity Section */
.activity-section {
    position: relative;
    z-index: 10;
}

.activity-placeholder {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 1rem;
    padding: 3rem;
    background: rgba(255, 255, 255, 0.02);
    border: 1px dashed rgba(255, 255, 255, 0.1);
    border-radius: 1rem;
    color: #6b7280;
}

.placeholder-icon {
    width: 2.5rem;
    height: 2.5rem;
    opacity: 0.5;
}

.activity-placeholder p {
    margin: 0;
    font-size: 0.9rem;
    font-family: monospace;
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
    opacity: 0.3;
}

.orb-1 {
    width: 500px;
    height: 500px;
    background: radial-gradient(circle, #6366f1 0%, transparent 70%);
    top: -150px;
    right: -100px;
}

.orb-2 {
    width: 400px;
    height: 400px;
    background: radial-gradient(circle, #8b5cf6 0%, transparent 70%);
    bottom: -100px;
    left: -100px;
}

/* Responsive Design */
@media (max-width: 1024px) {
    .stats-grid {
        grid-template-columns: repeat(2, 1fr);
    }
}

@media (max-width: 768px) {
    .home {
        padding: 1.5rem;
    }

    .welcome-section {
        padding: 1.5rem 0 2rem;
    }

    .actions-grid {
        grid-template-columns: 1fr;
    }

    .stat-card {
        padding: 1rem;
    }

    .stat-value {
        font-size: 1.25rem;
    }

    .orb-1 {
        width: 300px;
        height: 300px;
    }

    .orb-2 {
        width: 250px;
        height: 250px;
    }
}

@media (max-width: 480px) {
    .home {
        padding: 1rem;
        min-height: calc(100vh - 60px);
    }

    .stats-grid {
        grid-template-columns: repeat(2, 1fr);
        gap: 0.75rem;
    }

    .stat-card {
        flex-direction: column;
        text-align: center;
        padding: 1rem 0.75rem;
    }

    .stat-icon-wrapper {
        width: 2.5rem;
        height: 2.5rem;
    }

    .stat-icon {
        width: 1.25rem;
        height: 1.25rem;
    }

    .action-card {
        padding: 1.25rem;
    }

    .action-icon-wrapper {
        width: 3rem;
        height: 3rem;
    }

    .action-icon {
        width: 1.5rem;
        height: 1.5rem;
    }

    .welcome-subtitle {
        font-size: 1rem;
    }
}
</style>