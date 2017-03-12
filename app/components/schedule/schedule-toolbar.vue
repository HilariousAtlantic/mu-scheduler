<template>

  <div class="schedule-toolbar">

    <div class="schedule-options">

      <div class="btn-group">

        <button class="btn"><span class="fa fa-square"></span></button>

        <button class="btn"><span class="fa fa-th-large"></span></button>

      </div>

    </div>

    <div class="schedule-search">

      <button class="btn" @click="$emit('change', index-1)">

        <span class="fa fa-caret-left"></span>

      </button>

      <input
        type="text"
        class="form-control"
        :value="searchText"
        @focus="handleSearchFocus"
        @blur="focused = false"
        @keydown.enter="handleIndexChange">

      <button class="btn" @click="$emit('change', index+1)">

        <span class="fa fa-caret-right"></span>

      </button>

    </div>

  </div>

</template>

<script>

  export default {

    name: 'schedule-toolbar',

    props: ['index', 'length'],

    data() {

      return {

        focused: false

      }

    },

    computed: {

      searchText() {

        if (this.focused) {

          return this.index+1;

        } else return (this.index+1) + ' of ' + this.length;

      }

    },

    methods: {

      handleSearchFocus(event) {

        this.focused = true;

        setTimeout(() => event.target.select(), 10);

      },

      handleIndexChange(event) {

        let input = event.target;
        let index = parseInt(input.value)+1 || 0;

        this.$emit('change', index);
        input.blur();

      }

    }

  }

</script>

<style scoped>

  .schedule-toolbar {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .schedule-search button {
    padding: 6px 20px;
  }

  .schedule-search input {
    text-align: center;
  }

</style>
