<template>

  <div class="schedule-toolbar">

    <div class="schedule-options">

      <div class="btn-group tooltipped tooltipped-n"
        aria-label="Display options coming soon!">

        <button class="btn"><span class="fa fa-square"></span></button>

        <button class="btn"><span class="fa fa-th-large"></span></button>

      </div>

      <button class="btn" @click="showModal = true">Show CRNs</button>

      <modal :show="showModal" @close="showModal = false">

        <h3>Schedule CRNs</h3>

        <span v-for="crn in crns" class="crn">{{crn}}</span>

      </modal>

    </div>

    <div class="schedule-search">

      <button class="btn" @click="$emit('change', index-1)">

        <span class="fa fa-caret-left"></span>

      </button>

      <input
        type="text"
        class="form-control"
        :value="searchText"
        @focus="handleSearchFocus"
        @blur="focused = false"
        @keydown.enter="handleIndexChange">

      <button class="btn" @click="$emit('change', index+1)">

        <span class="fa fa-caret-right"></span>

      </button>

    </div>

  </div>

</template>

<script>

  import Modal from '../common/modal.vue';

  export default {

    name: 'schedule-toolbar',

    props: ['index', 'length'],

    components: {Modal},

    data() {

      return {

        focused: false,

        showModal: false

      }

    },

    computed: {

      searchText() {

        if (this.focused) {

          return this.index+1;

        } else return (this.index+1) + ' of ' + this.length;

      },

      crns() {

        let schedule = this.$store.getters.filteredSchedules[this.index];
        return schedule.courses.map(course => course.crn);

      }

    },

    methods: {

      handleSearchFocus(event) {

        this.focused = true;

        setTimeout(() => event.target.select(), 10);

      },

      handleIndexChange(event) {

        let input = event.target;
        let index = parseInt(input.value)-1 || 0;

        this.$emit('change', index);
        input.blur();

      }

    }

  }

</script>

<style scoped>

  .schedule-toolbar {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .schedule-search button {
    padding: 6px 20px;
  }

  .schedule-search input {
    text-align: center;
  }

  .modal h3 {
    margin-bottom: 10px;
  }

  .crn {
    height: 34px;
    padding: 7px 8px;
    font-size: 13px;
    color: #333;
    background-color: #fafafa;
    background-repeat: no-repeat;
    background-position: right 8px center;
    border: 1px solid #ddd;
    border-radius: 3px;
    outline: none;
    box-shadow: inset 0 1px 2px rgba(0,0,0,0.075);
  }

  .crn + .crn {
    margin-left: 5px;
  }

</style>
