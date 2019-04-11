<template>
  <sweet-modal title="Create a new ticket" ref="modal" class="text-left">
    <form
      @submit.prevent="createTicket(form)"
    >
        <span
            v-if="formError"
            v-html="formError"
            class="text-red mb-6 text-lg block"
          ></span>
        <div class="mb-4">
          <label
            class="block uppercase tracking-wide text-grey-darker text-xs font-bold mb-2"
            for="grid-first-name"
          >Ticket title</label>
          <input
            type="text"
            v-model="form.title"
            class="bg-grey-lighter appearance-none border-2 border-grey-lighter rounded w-full py-2 px-4 text-grey-darker leading-tight focus:outline-none focus:bg-white focus:border-purple"
          >
        </div>
        <div class="mb-4">
          <label
            class="block uppercase tracking-wide text-grey-darker text-xs font-bold mb-2"
            for="grid-first-name"
          >Ticket content</label>
          <textarea
            rows="6"
            v-model="form.content"
            class="bg-grey-lighter appearance-none border-2 border-grey-lighter rounded w-full py-2 px-4 text-grey-darker leading-tight focus:outline-none focus:bg-white focus:border-purple"
          ></textarea>
        </div>
      </div>
      <div class="text-right">
        <button
          @click.prevent="$refs.modal.close()"
          type="submit"
          class="focus:outline-none bg-white hover:bg-grey-lighter hover:text-grey-darker focus:outline-none text-grey font-bold py-3 px-4 rounded mr-4"
        >Cancel</button>
        <button
          type="submit"
          class="shadow bg-purple hover:bg-purple-light focus:shadow-outline focus:outline-none text-white font-bold py-3 px-4 rounded"
        >Submit ticket</button>
      </div>
    </form>
  </sweet-modal>
</template>

<script>
import { SweetModal } from "sweet-modal-vue";
export default {
  data: () => ({
    formError: null,
    form: {
      title: "",
      content: "",
      status: "waiting",
      author: ""
    }
  }),
  components: {
    SweetModal
  },
  created() {
    this.$bus.on("open-create-ticket-modal", this.openModal);
    this.$bus.on("close-create-ticket-modal", this.closeModal);
  },
  beforeDestroy() {
    this.$bus.off("open-create-ticket-modal", this.openModal);
    this.$bus.off("close-create-ticket-modal", this.closeModal);
  },
  methods: {
    openModal() {
      this.$refs.modal.open();
    },
    closeModal() {
      this.$refs.modal.close();
    },
    createTicket(ticket) {
      this.formError = null;
      this.form.author = this.$root.userName;
      if (
        this.form.title.length &&
        this.form.content.length &&
        this.form.status.length &&
        this.form.author.length
      ) {
        try {
          this.$bus.emit("add-ticket", ticket);
          this.closeModal();
        } catch {
          this.formError = "Something went wromg. Please try again.";
        }
      } else {
        this.formError = "Please fill in all fields.";
      }
    }
  }
};
</script>
<style>
textarea {
  resize: none;
}
.close {
  position: absolute;
  right: 0;
  top: 0;
  bottom: 0;
  display: flex;
  align-items: center;
  opacity: 0.3;
  width: 70px;
  text-align: center;
}
.close:hover {
  opacity: 1;
}
.close:before,
.close:after {
  position: absolute;
  left: 50%;
  content: " ";
  height: 33px;
  width: 2px;
  background-color: #333;
}
.close:before {
  transform: rotate(45deg);
}
.close:after {
  transform: rotate(-45deg);
}
</style>
