<template>

  <div class="filter-group" @click.stop>

    <button class="group-header btn btn-block" @click="toggleEditing">

      <span>{{label}}</span>

      <span v-if="editing" class="fa fa-caret-up"></span>
      <span v-else class="fa fa-caret-down"></span>

    </button>

    <div v-if="editing" class="group-editor">

      <slot></slot>

      <button class="create-button btn" @click="createFilter">

        <span class="fa fa-plus"></span>
        <span> Add Filter</span>

      </button>

    </div>

  </div>

</template>

<script>

  export default {

    name: 'filter-group',

    props: ['filterType', 'label'],

    computed: {

      editing() {

        return this.$store.state.selectedEditor == this.filterType;

      }

    },

    methods: {

      toggleEditing() {

        let type = this.editing ? '' : this.filterType;

        this.$store.dispatch('selectEditor', type);

      },

      createFilter() {

        this.$store.dispatch('createFilter', this.filterType);

      }

    }

  }

</script>

<style scoped>

  .filter-group {
    position: relative;
  }

  .group-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .group-editor {
    position: absolute;
    width: 520px;
    padding: 20px;
    background: #fff;
    border: 1px solid #ddd;
    border-radius: 3px;
    box-shadow: 0 3px 12px rgba(27,31,35,0.15);
    top: calc(100% + 5px);
    z-index: 10;
  }

</style>
