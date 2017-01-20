<template>

  <div class="time-filter">

    <schedule-filter :text="text" :active="filter.active" @edit="toggleEditing" @toggle="toggleActive"></schedule-filter>

    <modal v-if="showEditor">

      <filter-editor :result="previewText" @done="submitChanges" @quit="discardChanges">

        <time-filter-options :options="preview" @change="changePreview"></time-filter-options>

      </filter-editor>

    </modal>

  </div>

</template>

<script>

  import ScheduleFilter from './schedule-filter.vue';
  import Modal from '../common/modal.vue';
  import FilterEditor from './filter-editor.vue';
  import TimeFilterOptions from './time-filter-options.vue';

  import {toTime} from '../../lib/time';
  import {formatDayList} from '../../lib/days';

  export default {

    name: 'time-filter',

    props: ['filter'],

    components: {ScheduleFilter, Modal, FilterEditor, TimeFilterOptions},

    data() {

      return {

        showEditor: false,

        preview: Object.assign({}, this.filter.options)

      }

    },

    computed: {

      text() {

        let operator = this.filter.options.operator.toLowerCase();

        let time = toTime(this.filter.options.time);

        let days = formatDayList(this.filter.options.days);

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

        this.$store.dispatch('toggleScheduleFilter', this.filter.id);

      },

      changePreview(changes) {

        Object.assign(this.preview, changes);

      },

      submitChanges() {

        let id = this.filter.id;
        let changes = this.preview;

        this.$store.dispatch('updateScheduleFilter', {id, changes});
        this.toggleEditing();

      },

      discardChanges() {

        Object.assign(this.preview, this.filter.options);
        this.toggleEditing();

      }

    }

  }

</script>

<style scoped>

  .time-filter {

  }

</style>
