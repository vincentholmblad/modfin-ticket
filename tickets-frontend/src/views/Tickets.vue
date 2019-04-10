<template>
  <div class="">
    <app-header class="flex-no-grow" id="app-nav" v-model="search" />
    <tickets-list class="flex-grow" id="app-content" :filter="search" />
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
  },
  data: () => ({
    search: ''
  })
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
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  top: 60px;
}
</style>