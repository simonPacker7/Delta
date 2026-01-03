<script setup lang="ts">
import { computed } from 'vue'
import { UserCircleIcon } from '@heroicons/vue/24/solid'

interface Player {
    Name: string
    ID: string
}

const props = defineProps<{
    player1: Player | null
    player2: Player | null
    currentTurnId: string | null
    currentPlayerId: string | null
    isGameActive: boolean
}>()

const isPlayer1Turn = computed(() => props.currentTurnId === props.player1?.ID)
const isPlayer2Turn = computed(() => props.currentTurnId === props.player2?.ID)

const isCurrentPlayer = (playerId: string | undefined) => {
    return playerId === props.currentPlayerId
}
</script>

<template>
    <div class="game-header">
        <!-- Player 1 -->
        <div 
            class="player-card"
            :class="{ 
                'active-turn': isPlayer1Turn && isGameActive,
                'current-player': isCurrentPlayer(player1?.ID)
            }"
        >
            <div class="player-avatar">
                <UserCircleIcon class="avatar-icon" />
                <div v-if="isPlayer1Turn && isGameActive" class="turn-indicator"></div>
            </div>
            <div class="player-info">
                <span class="player-name">{{ player1?.Name || 'Player 1' }}</span>
                <span v-if="isCurrentPlayer(player1?.ID)" class="you-badge">You</span>
            </div>
        </div>

        <!-- VS Divider -->
        <div class="vs-divider">
            <span class="vs-text">VS</span>
        </div>

        <!-- Player 2 -->
        <div 
            class="player-card"
            :class="{ 
                'active-turn': isPlayer2Turn && isGameActive,
                'current-player': isCurrentPlayer(player2?.ID)
            }"
        >
            <div class="player-avatar">
                <UserCircleIcon class="avatar-icon" />
                <div v-if="isPlayer2Turn && isGameActive" class="turn-indicator"></div>
            </div>
            <div class="player-info">
                <span class="player-name">{{ player2?.Name || 'Player 2' }}</span>
                <span v-if="isCurrentPlayer(player2?.ID)" class="you-badge">You</span>
            </div>
        </div>
    </div>
</template>

<style scoped>
.game-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 1rem;
    padding: 1rem;
    background: rgba(255, 255, 255, 0.03);
    border: 1px solid rgba(255, 255, 255, 0.08);
    border-radius: 1rem;
    margin-bottom: 1rem;
}

.player-card {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    flex: 1;
    padding: 0.75rem;
    border-radius: 0.75rem;
    background: rgba(255, 255, 255, 0.02);
    border: 1px solid transparent;
    transition: all 0.3s ease;
}

.player-card:last-child {
    flex-direction: row-reverse;
    text-align: right;
}

.player-card:last-child .player-info {
    align-items: flex-end;
}

.player-card.active-turn {
    background: rgba(34, 197, 94, 0.1);
    border-color: rgba(34, 197, 94, 0.3);
}

.player-card.current-player {
    background: rgba(99, 102, 241, 0.1);
    border-color: rgba(99, 102, 241, 0.2);
}

.player-card.current-player.active-turn {
    background: linear-gradient(135deg, rgba(99, 102, 241, 0.15) 0%, rgba(34, 197, 94, 0.15) 100%);
    border-color: rgba(34, 197, 94, 0.4);
}

.player-avatar {
    position: relative;
    flex-shrink: 0;
}

.avatar-icon {
    width: 2.5rem;
    height: 2.5rem;
    color: #6b7280;
    transition: color 0.3s ease;
}

.player-card.active-turn .avatar-icon {
    color: #22c55e;
}

.player-card.current-player .avatar-icon {
    color: #6366f1;
}

.turn-indicator {
    position: absolute;
    bottom: 0;
    right: 0;
    width: 0.75rem;
    height: 0.75rem;
    background: #22c55e;
    border-radius: 50%;
    border: 2px solid #282828;
    animation: pulse-glow 2s ease-in-out infinite;
}

@keyframes pulse-glow {
    0%, 100% {
        box-shadow: 0 0 0 0 rgba(34, 197, 94, 0.5);
    }
    50% {
        box-shadow: 0 0 0 6px rgba(34, 197, 94, 0);
    }
}

.player-info {
    display: flex;
    flex-direction: column;
    gap: 0.25rem;
    min-width: 0;
}

.player-name {
    font-family: monospace;
    font-size: 0.95rem;
    font-weight: 600;
    color: #e5e7eb;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}

.you-badge {
    font-size: 0.7rem;
    font-weight: 600;
    text-transform: uppercase;
    letter-spacing: 0.05em;
    color: #a5b4fc;
    font-family: monospace;
}

.vs-divider {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
}

.vs-text {
    font-family: monospace;
    font-size: 0.8rem;
    font-weight: 700;
    color: #4b5563;
    padding: 0.375rem 0.75rem;
    background: rgba(255, 255, 255, 0.05);
    border-radius: 0.375rem;
}

/* Responsive */
@media (max-width: 480px) {
    .game-header {
        padding: 0.75rem;
        gap: 0.5rem;
    }

    .player-card {
        padding: 0.5rem;
        gap: 0.5rem;
    }

    .avatar-icon {
        width: 2rem;
        height: 2rem;
    }

    .player-name {
        font-size: 0.85rem;
    }

    .you-badge {
        font-size: 0.6rem;
    }

    .vs-text {
        font-size: 0.7rem;
        padding: 0.25rem 0.5rem;
    }
}
</style>
