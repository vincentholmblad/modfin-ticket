<template>
  <div class="flex w-full bg-purple-lightest">
    <ticket-list-column class="w-1/3 border-r" header="Waiting" slug="waiting"/>
    <ticket-list-column class="w-1/3 border-r" header="In Progress" slug="doing"/>
    <ticket-list-column class="w-1/3" header="Completed" slug="done"/>
    <sweet-modal title="Are you sure?" ref="deleteModal" class="text-left" width="400">
      <button
        @click="deleteTicketFlush()"
        class="shadow bg-purple hover:bg-purple-light focus:shadow-outline focus:outline-none text-white font-bold py-3 px-4 rounded w-full"
      >Yes, delete the ticket</button>
    </sweet-modal>
  </div>
</template>

<script>
import { SweetModal } from "sweet-modal-vue";
import TicketListColumn from "@/components/TicketListColumn.vue";
export default {
  props: ["filter"],
  data: () => ({
    tickets: [],
    deleteCheckTicket: null
  }),
  components: {
    TicketListColumn,
    SweetModal
  },
  mounted() {
    this.$store.dispatch("tickets/getAllTickets");
  },
  created() {
    this.$bus.on("change-status", this.changeStatus);
    this.$bus.on("add-ticket", this.addTicket);
    this.$bus.on("delete-ticket", this.deleteTicket);
  },
  beforeDestroy() {
    this.$bus.off("change-status", this.changeStatus);
    this.$bus.off("add-ticket", this.addTicket);
    this.$bus.off("delete-ticket", this.deleteTicket);
  },
  methods: {
    changeStatus(data) {
      var ticket = data.ticket;
      var status = data.status;

      if (ticket.status == status) {
        return;
      }

      this.$store.dispatch("tickets/updateTicketStatus", {
        ticket: ticket,
        status: status
      });
    },
    addTicket(ticket) {
      this.$store.dispatch("tickets/persistTicket", ticket);
    },
    deleteTicket(ticket) {
      this.$refs.deleteModal.open();
      this.deleteCheckTicket = ticket;
    },
    deleteTicketFlush() {
      this.$store.dispatch('tickets/removeTicket', this.deleteCheckTicket);
      this.deleteCheckTicket = null;
      this.$refs.deleteModal.close();
    }
  }
};
</script>

<style>
</style>
