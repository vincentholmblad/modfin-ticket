<template>
    <drag :transfer-data="ticket">
        <div class="item bg-white shadow rounded p-6 mb-4 cursor-move relative">
            <h1 class="text-xl mb-3 -mt-1">{{ ticket.title }}</h1>
            <p class="mb-4">{{ ticket.content }}</p>
            <p class="text-grey-dark text-sm">{{ ticket.author }} · {{ ticket.created_at | moment("from", "now") }}</p>
            <div class="absolute pin-r pin-t">
                <button class="delete focus:outline-none" @click="deleteTicket"></button>
            </div>
        </div>
    </drag>
</template>

<script>
export default {
    props: ['ticket'],
    methods: {
        deleteTicket() {
            this.$bus.emit('delete-ticket', this.ticket);
        }
    }
}
</script>

<style>
.delete {
  opacity: 0;
  position: relative;
  display: flex;
  align-items: center;
  width: 40px;
  height: 40px;
  text-align: center;
}
.item:hover .delete:hover {
  opacity: 1;
}
.delete:before,
.delete:after {
  position: absolute;
  left: 50%;
  content: " ";
  height: 20px;
  width: 2px;
  background-color: #333;
}
.delete:before {
  transform: rotate(45deg);
}
.delete:after {
  transform: rotate(-45deg);
}
.item:hover .delete {
    opacity: 0.3;
}
</style>
