<template>
<drop @drop="droppedTicket">
    <div class="w-full h-full overflow-y-auto text-left relative list-col">
        <div class="list-col-header px-6 py-4 border-b flex items-center justify-between sticky bg-purple-lightest">
            <h1 class="text-2xl leading-narrow">{{ header }}</h1>
            <span class="float-right font-thin text-base opacity-50">{{ tickets.length }} ticket(s)</span>
        </div>
        <div class="list-inner">
            <ticket-list-item v-for="(ticket, index) in tickets" v-bind:key="ticket.ticked_id && ticket.modified_at && index" :ticket="ticket" />
        </div>
    </div>
</drop>
</template>

<script>
import { mapGetters } from "vuex";
import TicketListItem from '@/components/TicketListItem.vue'
export default {
    props: ['header', 'slug'],
    components: {
        TicketListItem
    },
    computed: {
        tickets() {
            return this.$store.getters['tickets/ticketsByStatus'](this.slug);
        }
    },
    methods: {
        droppedTicket(data) {
            this.$bus.emit('change-status', { ticket: data, status: this.slug });
        }
    }
}
</script>

<style>
.list-col {
    transform: translateZ(0);
}
.list-inner {
    @apply px-6 py-4;
}
.list-col-header {
    top: 0;
    left: 0;
    right: 0;
}
</style>
