<template>

  <div class="time-filter">

    <schedule-filter :text="text" :active="active" @edit="toggleEditing" @toggle="toggleActive"></schedule-filter>

    <modal v-if="showEditor">

      <filter-editor :result="previewText" @done="submitChanges" @quit="discardChanges">

        <time-filter-editor :options="preview" @change="changePreview"></time-filter-editor>

      </filter-editor>

    </modal>

  </div>

</template>

<script>

  import ScheduleFilter from './schedule-filter.vue';
  import Modal from '../modal.vue';
  import FilterEditor from './filter-editor.vue';
  import TimeFilterEditor from './time-filter-editor.vue';

  import {toTime} from '../../lib/time';
  import {formatDayList} from '../../lib/days';

  export default {

    name: 'time-filter',

    props: ['options', 'active'],

    components: {ScheduleFilter, Modal, FilterEditor, TimeFilterEditor},

    data() {

      return {

        showEditor: false,

        preview: Object.assign({}, this.options)

      }

    },

    computed: {

      text() {

        let operator = this.options.operator.toLowerCase();

        let time = toTime(this.options.time);

        let days = formatDayList(this.options.days);

        return 'I want to ' + operator + ' ' + time + ' on ' + days;

      },

      previewText() {

        let operator = this.preview.operator.toLowerCase();

        let time = toTime(this.preview.time);

        let days = formatDayList(this.preview.days);

        return 'I want to ' + operator + ' ' + time + ' on ' + days;

      }

    },

    methods: {

      toggleEditing() {

        this.showEditor = !this.showEditor;

      },

      toggleActive() {

        // TODO dispatch a filter change action

      },

      changePreview(changes) {

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

</style>
