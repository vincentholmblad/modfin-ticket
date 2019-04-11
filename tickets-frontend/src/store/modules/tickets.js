import ticketApi from '../../api/ticket'
import Vue from 'vue'

// initial state
const state = {
    all: [],
    filter: ''
}

// getters
const getters = {
    ticketsByStatus: (state) => (status) => {
        return state.all.filter(ticket => {
            return ticket.status.toLowerCase() == status && (
                ticket.title.toLowerCase().includes(state.filter.toLowerCase()) ||
                ticket.content.toLowerCase().includes(state.filter.toLowerCase()) ||
                ticket.author.toLowerCase().includes(state.filter.toLowerCase())
            )
        }).sort((a, b) => {
            return (new Date(b.modified_at)).getTime() - (new Date(a.modified_at)).getTime()
        })
    }
}

// actions
const actions = {
    getAllTickets({ commit }) {
        ticketApi.getTickets().then(tickets => {
            commit('setTickets', tickets)
        })
    },
    persistTicket({ commit }, ticket) {
        ticketApi.persistTicket(ticket).then(persistedTicket => {
            commit('addTicket', persistedTicket)
        })
    },
    updateTicketStatus({ commit }, data) {
        ticketApi.updateTicketStatus(data.ticket.ticket_id, data.status).then(ticket => {
            commit('updateTicket', ticket)
        })
    },
    removeTicket({ commit }, ticket) {
        ticketApi.removeTicket(ticket.ticket_id).then(success => {
            commit('removeTicket', ticket)
        })
    },
    filter({ commit }, filter) {
        commit('setFilter', filter);
    }
}

// mutations
const mutations = {
    setTickets(state, tickets) {
        state.all = tickets
    },
    addTicket(state, ticket) {
        state.all.push(ticket);
    },
    updateTicket(state, ticket) {
        var ticketIndex = state.all.map(function (e) { return e.ticket_id; }).indexOf(ticket.ticket_id);
        Vue.set(state.all, ticketIndex, ticket);
    },
    removeTicket(state, ticket) {
        var ticketIndex = state.all.map(function (e) { return e.ticket_id; }).indexOf(ticket.ticket_id);
        Vue.delete(state.all, ticketIndex);
    },
    setFilter(state, filter) {
        state.filter = filter;
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}