<template>
  <div class="">
    <app-header class="flex-no-grow" id="app-nav" />
    <tickets-list class="flex-grow" id="app-content" />
    <create-ticket-modal />
  </div>
</template>

<script>
import AppHeader from '@/components/Header.vue'
import CreateTicketModal from '@/components/CreateTicketModal.vue'
import TicketsList from '@/components/TicketsList.vue'

export default {
  name: 'tickets',
  beforeRouteEnter (to, from, next) {
    if(!from.name) {
      next('/');
    }
    next();
  },
  components: {
    AppHeader,
    CreateTicketModal,
    TicketsList
  }
}
</script>

<style>
#app-nav {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 2;
  height: 60px;
}
#app-content {
  width: 100vw;
  height: 100vh;
  padding-top: 60px;
  overflow: hidden;
}
</style>