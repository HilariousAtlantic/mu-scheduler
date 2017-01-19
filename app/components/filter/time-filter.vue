<template>

  <div class="time-filter">

    <schedule-filter :text="text" :active="active" @edit="toggleEditor"></schedule-filter>

    <div v-if="showEditor" class="modal">

      <filter-editor :result="text">

        <div class="filter-options">

          <input type="text">

        </div>

      </filter-editor>

    </div>

  </div>

</template>

<script>

  import ScheduleFilter from './schedule-filter.vue';
  import FilterEditor from './filter-editor.vue';

  export default {

    name: 'time-filter',

    props: ['options', 'active'],

    components: {ScheduleFilter, FilterEditor},

    data() {

      return {

        showEditor: false

      }

    },

    computed: {

      text() {

        return 'I want to ' + this.options.operator.toLowerCase() + ' ' + this.options.time + ' on ' + this.options.days;

      }

    },

    methods: {

      toggleEditor() {

        this.showEditor = !this.showEditor;

      }

    }

  }

</script>

<style scoped>

  .time-filter {

  }

  .modal {

    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    background: rgba(0, 0, 0, .5);
    z-index: 1;

  }

</style>
