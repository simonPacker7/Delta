<script setup lang="ts">
const props = defineProps({
    id: String,
    title: String,
    active: Boolean,
    showCloseBtn: {
        type: Boolean,
        default: true,
    },
})

</script>

<template>
    <div>
        <section :id="props.id" ref="modalRef" class="modal" :class="{ 'modal-active': active }" role="dialog">
            <div class="modal-content modal-dialog">
                <slot name="header">
                    <header class="modal-header center">
                        <h2 class="modal-title">
                            {{ props.title }}
                        </h2>
                        <button v-if="showCloseBtn" type="button" class="close close-btn" data-dismiss="modal"
                            aria-label="Close">
                            <span aria-hidden="true">&times;</span>
                        </button>
                    </header>
                </slot>
                <slot name="body">
                    <div class="modal-body">
                        <slot name="content">
                        </slot>
                    </div>
                </slot>
                <slot name="footer">
                    <div class="modal-footer">
                        <slot name="footerContent">
                            <button class="btn btn-outline-primary">Exit</button>
                        </slot>
                    </div>
                </slot>
            </div>
        </section>
    </div>
</template>

<style scoped>
.modal-active {
    display: block !important;
}

.modal {
    display: none;
    /* Hidden by default */
    position: fixed;
    z-index: 1;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    overflow: auto;
    background-color: rgba(0, 0, 0, 0.4)
}

.modal-content {
    margin: 15% auto;
    border: 1px solid #888;
    width: 80%;
}

.close-btn {
    border: 0px;
}

.modal-header,
.modal-body,
.modal-footer {
    background-color: var(--vt-c-black-soft);
}

.modal-header {
    border-bottom: 1px dashed beige;
}

.modal-footer {
    border-top: 0px;
}

.modal-title {
    font-family: monospace;
    font-weight: 500;
    color: beige;
}

.center {
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>