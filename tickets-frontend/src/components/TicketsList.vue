<template>
    <div class="flex w-full bg-purple-lightest">
        <ticket-list-column :tickets="waitingTickets" class="w-1/3 border-r" header="Waiting" />
        <ticket-list-column :tickets="doingTickets" class="w-1/3 border-r" header="Doing" />
        <ticket-list-column :tickets="doneTickets" class="w-1/3" header="Done" />
    </div>
</template>

<script>
import TicketListColumn from '@/components/TicketListColumn.vue'
export default {
    props: [
        'filter'
    ],
    data: () => ({
        tickets: []
    }),
    components: {
        TicketListColumn
    },
    mounted() {
        this.getTickets()
    },
    created() {
        this.$bus.on('change-status', this.changeStatus);
        this.$bus.on("add-ticket", this.addTicket);
    },
    beforeDestroy() {
        this.$bus.off('change-status', this.changeStatus);
        this.$bus.off("add-ticket", this.addTicket);
    },
    computed: {
        waitingTickets() {
            return this.tickets.filter(ticket => {
                return ticket.status.toLowerCase() == 'waiting' && (ticket.title.toLowerCase().includes(this.filter.toLowerCase()) || ticket.content.toLowerCase().includes(this.filter.toLowerCase()) || ticket.author.toLowerCase().includes(this.filter.toLowerCase())) 
            })
        },
        doingTickets() {
            return this.tickets.filter(ticket => {
                return ticket.status.toLowerCase() == 'doing' && (ticket.title.toLowerCase().includes(this.filter.toLowerCase()) || ticket.content.toLowerCase().includes(this.filter.toLowerCase()) || ticket.author.toLowerCase().includes(this.filter.toLowerCase())) 
            })
        },
        doneTickets() {
            return this.tickets.filter(ticket => {
                return ticket.status.toLowerCase() == 'done' && (ticket.title.toLowerCase().includes(this.filter.toLowerCase()) || ticket.content.toLowerCase().includes(this.filter.toLowerCase()) || ticket.author.toLowerCase().includes(this.filter.toLowerCase())) 
            })
        }
    },
    methods: {
        getTickets() {
            this.axios.get('http://localhost:8000/api/tickets/')
            .then(data => {
                this.tickets = data.data
            })
            .catch(e => {
                console.log(e)
            })
        },
        changeStatus(data) {
            var ticket = data.ticket;
            var status = data.status;

            if(ticket.status == status) {
                return;
            }

            var tempStatus = ticket.status;

            // Nececarry for reactivity in computed tickets for columns
            ticket.status = status;
            ticket.modified_at = new Date();

            this.axios.post('http://localhost:8000/api/tickets/update', {
                ticket_id: ticket.ticket_id,
                status: status
            })
            .then(data => {
                this.$set(this.tickets, [this.tickets.indexOf(data.ticket)], data.data);
            })
            .catch(e => {
                console.log(e);
                this.$set(this.tickets[this.tickets.indexOf(data.ticket)], 'status', tempStatus);
            })
        },
        addTicket(ticket) {
            this.tickets.push(ticket);
        }
    }
}
</script>

<style>

</style>
