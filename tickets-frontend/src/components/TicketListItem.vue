<template>
    <drag :transfer-data="ticket">
        <div class="bg-white shadow rounded p-6 mb-4 cursor-move">
            <h1 class="text-xl mb-4 -mt-1">{{ ticket.title }}</h1>
            <p class="mb-4">{{ ticket.content }}</p>
            <p class="text-grey-dark text-sm">{{ ticket.author }} · {{ ticket.created_at | moment("from", "now") }}</p>
        </div>
    </drag>
</template>

<script>
export default {
    props: ['ticket'],
    methods: {
        updateStatus(status) {
            this.axios.post('http://localhost:8000/api/tickets/' + ticket.ticket_id, {
                status: status
            })
            .then(data => {
                console.log(data)
            })
            .catch(e => {
                console.log(e)
            })
        }
    }
}
</script>

<style>

</style>
