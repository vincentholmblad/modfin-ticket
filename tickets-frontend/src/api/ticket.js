import Vue from 'vue'
export default {
    async getTickets() {
        var tickets = await new Promise((resolve, reject) => {
            Vue.axios
                .get('http://localhost:8000/api/tickets/')
                .then(data => {
                    resolve(data.data);
                })
                .catch(e => {
                    reject(e);
                })
        })
        return tickets;
    },
    async persistTicket(ticket) {
        var persistedTicket = await new Promise((resolve, reject) => {
            Vue.axios
                .post("http://localhost:8000/api/tickets/", ticket)
                .then(data => {
                    resolve(data.data);
                })
                .catch(e => {
                    reject(e);
                });
        })
        return persistedTicket;
    },
    async updateTicketStatus(ticketId, status) {
        var updatedTicket = await new Promise((resolve, reject) => {
            Vue.axios
                .post("http://localhost:8000/api/tickets/update", {
                    ticket_id: ticketId,
                    status: status
                })
                .then(data => {
                    resolve(data.data);
                })
                .catch(e => {
                    reject(e);
                });
        })
        return updatedTicket;
    },
    async removeTicket(ticketId) {
        var deleted = await new Promise((resolve, reject) => {
            Vue.axios
                .post("http://localhost:8000/api/tickets/delete", {
                    ticket_id: ticketId
                })
                .then(data => {
                    resolve(data.data);
                })
                .catch(e => {
                    reject(e);
                });
        })
        return deleted == 'deleted';
    }
}