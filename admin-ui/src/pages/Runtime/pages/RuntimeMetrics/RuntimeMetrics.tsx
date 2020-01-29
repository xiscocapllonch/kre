import React, { useState } from 'react';

import { Row, RowsWrapper } from 'react-grid-resizable';
import DashboardTitle from './components/DashboardTitle/DashboardTitle';

import GeneralInfo from './boxes/GeneralInfo/GeneralInfo';
import Accuracy from './boxes/Accuracy/Accuracy';
import LabelStats from './boxes/LabelStats/LabelStats';
import ConfusionMatrixBox from './boxes/ConfusionMatrixBox/ConfusionMatrixBox';

import dataSimple from './data_simple.json';
import dataComplete from './data_complete.json';

import cx from 'classnames';
import styles from './RuntimeMetrics.module.scss';

function RuntimeMetrics() {
  const separatorRowProps = { className: styles.separatorRow };

  const [expanded, setExpanded] = useState<string>('');

  // TODO: check version
  const data = true ? dataSimple : dataComplete;

  function toggleExpanded(nodeId: string): void {
    if (expanded) {
      setExpanded('');
    } else {
      setExpanded(nodeId);
    }
  }

  function getNodesToExpand() {
    const nodes = [expanded];
    let act = expanded;

    while (act.length > 0) {
      act = act.slice(0, -2);
      nodes.push(act);
    }

    return nodes;
  }

  const minimize = {
    [styles.minimize]: expanded
  };

  const nodesToExpand = getNodesToExpand();

  const height = expanded ? window.innerHeight - 164 : '100%';
  const width = window.innerWidth - 310;

  const SuccessFailsHeight = width / 4;

  const nLabels = Math.sqrt(data.confusionMatrix.length);
  const confusionMatrixHeight = nLabels * 100;
  const SeriesHeight = nLabels * 100;

  return (
    <div className={styles.container}>
      <DashboardTitle runtimeName={'Runtime X'} versionName={'V1.0.2'} />
      <div className={styles.content}>
        <div
          className={cx(styles.wrapper, {
            [styles.expanded]: expanded
          })}
          style={{ height }}
        >
          <RowsWrapper separatorProps={separatorRowProps}>
            <Row
              initialHeight={165}
              style={{
                maxHeight: '165px'
              }}
              className={cx(styles.row, minimize, {
                [styles.maximize]: nodesToExpand.includes('r1')
              })}
              disabled
            >
              <GeneralInfo data={data.general} />
            </Row>
            <Row
              initialHeight={confusionMatrixHeight}
              style={{
                maxHeight: `${confusionMatrixHeight}px`,
                marginTop: '25px'
              }}
              className={cx(styles.row, minimize, {
                [styles.maximize]: nodesToExpand.includes('r2')
              })}
              top={false}
            >
              <ConfusionMatrixBox
                toggleExpanded={toggleExpanded}
                nodeId={'r2'}
                data={data.confusionMatrix}
              />
            </Row>
            <Row
              initialHeight={SeriesHeight}
              className={cx(styles.row, minimize, {
                [styles.maximize]: nodesToExpand.includes('r3')
              })}
              style={{ maxHeight: 277 + 590 - 160 }}
            >
              <LabelStats
                toggleExpanded={toggleExpanded}
                nodeId={'r3'}
                data={data.series}
              />
            </Row>
            <Row
              initialHeight={SuccessFailsHeight}
              className={cx(styles.row, minimize, {
                [styles.maximize]: nodesToExpand.includes('r4')
              })}
              style={{ maxHeight: 277 + 590 - 160 }}
            >
              <Accuracy
                toggleExpanded={toggleExpanded}
                nodeId={'r4'}
                data={data.successVsFails}
              />
            </Row>
          </RowsWrapper>
        </div>
      </div>
    </div>
  );
}

export default RuntimeMetrics;
