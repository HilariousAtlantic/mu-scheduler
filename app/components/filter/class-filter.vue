<template>

  <div class="class-filter">

    <schedule-filter :text="text" :active="filter.active" @edit="toggleEditing" @toggle="toggleActive"></schedule-filter>

    <modal v-if="showEditor">

      <filter-editor :result="previewText" @done="submitChanges" @quit="discardChanges">

        <class-filter-options :options="preview" @change="changePreview"></class-filter-options>

      </filter-editor>

    </modal>

  </div>

</template>

<script>

  import ScheduleFilter from './schedule-filter.vue';
  import Modal from '../common/modal.vue';
  import FilterEditor from './filter-editor.vue';
  import ClassFilterOptions from './class-filter-options.vue';

  import {formatDayList} from '../../lib/days';

  export default {

    name: 'class-filter',

    props: ['filter'],

    components: {ScheduleFilter, Modal, FilterEditor, ClassFilterOptions},

    data() {

      return {

        showEditor: false,

        preview: Object.assign({}, this.filter.options)

      }

    },

    computed: {

      text() {

        let operator = this.filter.options.operator.toLowerCase();

        let amount = this.filter.options.amount;

        let days = formatDayList(this.filter.options.days);

        return 'I want ' + operator + ' ' + amount + ' classes on ' + days;

      },

      previewText() {

        let operator = this.preview.operator.toLowerCase();

        let amount = this.preview.amount;

        let days = formatDayList(this.preview.days);

        return 'I want ' + operator + ' ' + amount + ' classes on ' + days;

      }

    },

    methods: {

      toggleEditing() {

        this.showEditor = !this.showEditor;

      },

      toggleActive() {

        this.$store.dispatch('updateScheduleFilter', {id: this.filter.id, changes: {active: !this.filter.active}});

      },

      changePreview(changes) {

        Object.assign(this.preview, changes);

      },

      submitChanges() {

        let {operator, amount, days} = this.preview;

        this.$store.dispatch('updateScheduleFilter', {id: this.filter.id, changes: {options: {operator, amount, days}}});
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

  .class-filter {

  }

</style>
