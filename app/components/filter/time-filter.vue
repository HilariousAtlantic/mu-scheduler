<template>

  <div class="time-filter">

    <schedule-filter :text="text" :active="active" @edit="toggleEditing"></schedule-filter>

    <div v-if="showEditor" class="modal">

      <filter-editor :result="previewText" @done="submitChanges" @quit="discardChanges">

        <time-filter-editor :options="preview" @change="onPreviewChange"></time-filter-editor>

      </filter-editor>

    </div>

  </div>

</template>

<script>

  import ScheduleFilter from './schedule-filter.vue';
  import FilterEditor from './filter-editor.vue';
  import TimeFilterEditor from './time-filter-editor.vue';

  import {toTime} from '../../lib/time';

  export default {

    name: 'time-filter',

    props: ['options', 'active'],

    components: {ScheduleFilter, FilterEditor, TimeFilterEditor},

    data() {

      return {

        showEditor: false,

        preview: Object.assign({}, this.options)

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

        let time = toTime(this.options.time);

        let days = this.options.days.map(day => abbreviations[day]).join(', ');

        return 'I want to ' + operator + ' ' + time + ' on ' + days;

      },

      previewText() {

        let abbreviations = {

          M: 'Monday',
          T: 'Tuesday',
          W: 'Wednesday',
          R: 'Thursday',
          F: 'Friday',
          S: 'Saturday'

        }

        let operator = this.preview.operator.toLowerCase();

        let time = toTime(this.preview.time);

        let days = this.preview.days.map(day => abbreviations[day]).join(', ');

        return 'I want to ' + operator + ' ' + time + ' on ' + days;

      }

    },

    methods: {

      toggleEditing() {

        this.showEditor = !this.showEditor;

      },

      onPreviewChange(changes) {

        Object.assign(this.preview, changes);

      },

      submitChanges() {

        // TODO dispatch a filter change action
        this.toggleEditing();

      },

      discardChanges() {

        Object.assign(this.preview, this.options);
        this.toggleEditing();

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
