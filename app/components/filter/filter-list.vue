<template>

  <div class="filter-list">

    <filter-group
      label="Start Time"
      :editing="editing == 'start'"
      @edit="selectEditor('start')">

      <start-filter
        v-for="filter in getFilters('start')"
        :id="filter.id"
        :options="filter.options"
      ></start-filter>

      <button class="create-button btn" @click="$store.dispatch('createFilter', 'start')">

        <span class="fa fa-plus"></span>
        <span> Add Filter</span>

      </button>

    </filter-group>

    <filter-group
      label="Finish Time"
      :editing="editing == 'finish'"
      @edit="selectEditor('finish')">

      <finish-filter
        v-for="filter in getFilters('finish')"
        :id="filter.id"
        :options="filter.options"
      ></finish-filter>

      <button class="create-button btn" @click="$store.dispatch('createFilter', 'finish')">

        <span class="fa fa-plus"></span>
        <span> Add Filter</span>

      </button>

    </filter-group>

    <filter-group
      label="Break Time"
      :editing="editing == 'break'"
      @edit="selectEditor('break')">

      <break-filter
        v-for="filter in getFilters('break')"
        :id="filter.id"
        :options="filter.options"
      ></break-filter>

      <button class="create-button btn" @click="$store.dispatch('createFilter', 'break')">

        <span class="fa fa-plus"></span>
        <span> Add Filter</span>

      </button>

    </filter-group>

    <filter-group
      label="Class Load"
      :editing="editing == 'class'"
      @edit="selectEditor('class')">

      <class-filter
        v-for="filter in getFilters('class')"
        :id="filter.id"
        :options="filter.options"
      ></class-filter>

      <button class="create-button btn" @click="$store.dispatch('createFilter', 'class')">

        <span class="fa fa-plus"></span>
        <span> Add Filter</span>

      </button>

    </filter-group>

    <button
      class="btn btn-danger"
      :class="{disabled: !canClearFilters}"
      @click="clearFilters">

      <span>Remove Filters</span>

    </button>

    <button
      class="btn btn-primary"
      :class="{disabled: !canApplyFilters}"
      @click="applyFilters">

      <span>Apply Changes</span>

    </button>

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

    computed: {

      canClearFilters() {

        return this.$store.state.filters.length > 0;

      },

      canApplyFilters() {

        for (let filter of this.$store.state.filters) {

          if (filter.changes) return true;

        }

        return false;

      }

    },

    methods: {

      selectEditor(type) {

        this.editing = this.editing == type ? '' : type;

      },

      getFilters(type) {

        return this.$store.state.filters.filter(filter => filter.type === type);

      },

      clearFilters() {

        if (this.canClearFilters) {

          this.$store.dispatch('clearFilters');

        }

      },

      applyFilters() {

        if (this.canApplyFilters) {

          this.$store.dispatch('applyFilters');

        }

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

  .filter-group + .filter-group, .btn {
    margin-left: 5px;
  }

  .filter {
    margin-bottom: 10px;
  }

</style>
