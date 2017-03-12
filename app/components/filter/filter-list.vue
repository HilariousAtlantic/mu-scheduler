<template>

  <div class="filter-list">

    <filter-group
      label="Start Time"
      :index="0/4"
      :editing="editing == 'start'"
      @edit="selectEditor('start')">

      <start-filter
        v-for="filter in getFilters('start')"
        :id="filter.id"
        :active="filter.active"
        :options="filter.options"
      ></start-filter>

      <button class="create-button btn" @click="$store.dispatch('createFilter', 'start')">

        <span class="fa fa-plus"></span>
        <span> Add Filter</span>

      </button>

    </filter-group>

    <filter-group
      label="Finish Time"
      :index="1/4"
      :editing="editing == 'finish'"
      @edit="selectEditor('finish')">

      <finish-filter
        v-for="filter in getFilters('finish')"
        :active="filter.active"
        :options="filter.options"
      ></finish-filter>

      <button class="create-button btn" @click="$store.dispatch('createFilter', 'finish')">

        <span class="fa fa-plus"></span>
        <span> Add Filter</span>

      </button>

    </filter-group>

    <filter-group
      label="Break Time"
      :index="2/4"
      :editing="editing == 'break'"
      @edit="selectEditor('break')">

      <break-filter
        v-for="filter in getFilters('break')"
        :active="filter.active"
        :options="filter.options"
      ></break-filter>

      <button class="create-button btn" @click="$store.dispatch('createFilter', 'break')">

        <span class="fa fa-plus"></span>
        <span> Add Filter</span>

      </button>

    </filter-group>

    <filter-group
      label="Class Load"
      :index="3/4"
      :editing="editing == 'class'"
      @edit="selectEditor('class')">

      <class-filter
        v-for="filter in getFilters('class')"
        :active="filter.active"
        :options="filter.options"
      ></class-filter>

      <button class="create-button btn" @click="$store.dispatch('createFilter', 'class')">

        <span class="fa fa-plus"></span>
        <span> Add Filter</span>

      </button>

    </filter-group>

    <filter-group
      label="Sections"
      :index="4/4"
      :editing="editing == 'section'"
      @edit="selectEditor('section')">

      <section-filter></section-filter>

    </filter-group>

    <button
      class="apply-button btn btn-primary"
      @click="$store.dispatch('applyFilters')"
    >Apply Changes</button>

  </div>

</template>

<script>

  import StartFilter from './start-filter.vue';
  import FinishFilter from './finish-filter.vue';
  import BreakFilter from './break-filter.vue';
  import ClassFilter from './class-filter.vue';
  import SectionFilter from './section-filter.vue';
  import FilterGroup from './filter-group.vue';

  export default {

    name: 'filter-list',

    components: {StartFilter, FinishFilter, BreakFilter, ClassFilter, SectionFilter, FilterGroup},

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

  .filter-group + .filter-group, .apply-button {
    margin-left: 5px;
  }

  .create-button {
    margin-top: 10px;
  }

</style>
