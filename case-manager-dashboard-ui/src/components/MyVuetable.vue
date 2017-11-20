<template>
  <vuetable ref="vuetable"
    api-url="http://localhost:3000/active_applications"
    :fields="fields"
  ></vuetable>
</template>

<script>
import Vuetable from 'vuetable-2/src/components/Vuetable'

export default {
  components: {
    Vuetable
  },
  data () {
    return {
      fields: [
        {
          name: 'client_id',
          title: 'CLTC#'
        },
        {
          name: 'client.area.name',
          title: 'Area'
        },
        {
          name: 'client.first_name',
          title: 'First Name'
        },
        {
          name: 'client.last_name',
          title: 'Last Name'
        },
        {
          name: 'program.name',
          title: 'Program'
        },
        {
          name: 'case_management_activity_due_this_month',
          title: 'Activity Due'
        },
        {
          name: 'case_management_activity_due_this_month_date_due',
          title: 'Date Due'
        },
        {
          name: 'case_management_activity_due_this_month_date_done',
          title: 'Date Done'
        },
        {
          name: 'next_quarterly_visit_date_due',
          title: 'Next QV Due'
        },
        {
          name: 'next_re_evaluation_date_due',
          title: 'Next RE Due'
        },
      ]
    }
  }
}
</script>
