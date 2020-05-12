import React from 'react';
import Chip from '../../../../../../../components/Chip/Chip';
import { isEditable, ProcessChip } from './AppliedFilters';
import moment from 'moment';

const filterToLabel = new Map([
  ['level', 'Level'],
  ['search', 'Search'],
  ['processes', 'Processes'],
  ['nodeName', 'Process'],
  ['startDate', 'From'],
  ['endDate', 'To'],
  ['workflowId', 'Workflow']
]);
function getFilterLabel(filter: string) {
  return filterToLabel.get(filter);
}

function getValueLabel(filter: string, value: string | ProcessChip) {
  if (filter === 'startDate' || filter === 'endDate') {
    if (typeof value === 'string')
      return moment(value).format('MMM DD, YYYY HH:mm');
  }
  if (filter === 'processes') {
    if (typeof value !== 'string')
      return `${value.workflowName} - ${value.processes.join(', ')}`;
  }

  return value;
}

type Props = {
  filter: string;
  value: string;
  removeFilters: Function;
};
export function Filter({ filter, value, removeFilters }: Props) {
  const label = `${getFilterLabel(filter)}: ${getValueLabel(filter, value)}`;

  function removeFilter() {
    removeFilters({ [filter]: value });
  }
  return (
    <Chip
      label={label}
      onClose={isEditable(filter) ? removeFilter : undefined}
    />
  );
}

export default Filter;
