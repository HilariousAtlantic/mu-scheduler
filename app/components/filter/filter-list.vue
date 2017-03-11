<template>

  <div class="filter-list">

    <filter-group
      label="Start Time"
      filtertype="start"
      :index="0/4"
      :editing="editing == 'start'"
      @edit="selectEditor('start')">

      <start-filter
        v-for="filter in getFilters('start')"
        :active="filter.active"
        :options="filter.options"
      ></start-filter>

    </filter-group>

    <filter-group
      label="Finish Time"
      filtertype="finish"
      :index="1/4"
      :editing="editing == 'finish'"
      @edit="selectEditor('finish')">

      <finish-filter
        v-for="filter in getFilters('finish')"
        :active="filter.active"
        :options="filter.options"
      ></finish-filter>

    </filter-group>

    <filter-group
      label="Break Time"
      filtertype="break"
      :index="2/4"
      :editing="editing == 'break'"
      @edit="selectEditor('break')">

      <break-filter
        v-for="filter in getFilters('break')"
        :active="filter.active"
        :options="filter.options"
      ></break-filter>

    </filter-group>

    <filter-group
      label="Class Load"
      filtertype="class"
      :index="3/4"
      :editing="editing == 'class'"
      @edit="selectEditor('class')">

      <class-filter
        v-for="filter in getFilters('class')"
        :active="filter.active"
        :options="filter.options"
      ></class-filter>

    </filter-group>

    <filter-group
      label="Sections"
      :index="4/4"
      :editing="editing == 'section'"
      @edit="selectEditor('section')">

    </filter-group>

  </div>

</template>

<script>

  import StartFilter from './start-filter.vue';
  import FinishFilter from './finish-filter.vue';
  import BreakFilter from './break-filter.vue';
  import ClassFilter from './class-filter.vue';
  import FilterGroup from './filter-group.vue';

  export default {

    name: 'filter-list',

    components: {StartFilter, FinishFilter, BreakFilter, ClassFilter, FilterGroup},

    data() {

      return {

        editing: ''

      }

    },

    methods: {

      selectEditor(type) {

        this.editing = this.editing == type ? '' : type;

      },

      getFilters(type) {

        return this.$store.state.filters.filter(filter => filter.type === type);

      }

    }

  }

</script>

<style scoped>

  .filter-list {
    display: flex;
  }

  .filter-group {
    flex: 1;
  }

  .filter-group + .filter-group {
    margin-left: 5px;
  }

</style>
