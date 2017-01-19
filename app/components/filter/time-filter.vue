<template>

  <div class="time-filter">

    <schedule-filter :text="text" :active="active" @edit="toggleEditor"></schedule-filter>

    <div v-if="showEditor" class="modal">

      <filter-editor :result="text" @done="toggleEditor">

        <div class="filter-options">

          <dropdown
            :options="['Start Before', 'Start After', 'Finish Before', 'Finish After']"
            :defaultOption="options.operator"
            @select="handleOperatorChange"
          ></dropdown>

          <time-input
            :step="15"
            :defaultTime="options.time"
          ></time-input>

          <days-input
            :days="['M', 'T', 'W', 'R', 'F', 'S']"
            :defaultDays="options.days"
            @change="handleDaysChange"
          ></days-input>

        </div>

      </filter-editor>

    </div>

  </div>

</template>

<script>

  import ScheduleFilter from './schedule-filter.vue';
  import FilterEditor from './filter-editor.vue';
  import Dropdown from '../editor/dropdown.vue';
  import TimeInput from '../editor/time-input.vue';
  import DaysInput from '../editor/days-input.vue';

  export default {

    name: 'time-filter',

    props: ['options', 'active'],

    components: {ScheduleFilter, FilterEditor, Dropdown, TimeInput, DaysInput},

    data() {

      return {

        showEditor: false

      }

    },

    computed: {

      text() {

        let abbreviations = {

          M: 'Monday',
          T: 'Tuesday',
          W: 'Wednesday',
          R: 'Thursday',
          F: 'Friday',
          S: 'Saturday'

        }

        let operator = this.options.operator.toLowerCase();

        let days = this.options.days.map(day => abbreviations[day]).join(', ');

        return 'I want to ' + operator + ' ' + this.options.time + ' on ' + days;

      }

    },

    methods: {

      toggleEditor() {

        this.showEditor = !this.showEditor;

      },

      handleOperatorChange() {



      },

      handleDaysChange() {



      }

    }

  }

</script>

<style scoped>

  .time-filter {

  }

  .filter-options {

    display: inline-block;

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
