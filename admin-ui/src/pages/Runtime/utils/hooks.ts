import { useMutation } from '@apollo/react-hooks';
import {
  ACTIVATE_VERSION,
  DEACTIVATE_VERSION,
  DEPLOY_VERSION,
  STOP_VERSION,
  ActivateVersionResponse,
  DeactivateVersionResponse,
  DeployVersionResponse,
  StopVersionResponse,
  VersionActionVars
} from '../pages/RuntimeStatusPreview/RuntimeStatusPreview.graphql';

export enum actions {
  activate = 'activateVersion',
  deploy = 'deployVersion',
  stop = 'stopVersion',
  deactivate = 'deactivateVersion'
}

export default function useVersionAction(onCompleted: any = () => {}) {
  const [activateMutation, { loading: loadingM1 }] = useMutation<
    ActivateVersionResponse,
    VersionActionVars
  >(ACTIVATE_VERSION, { onCompleted });
  const [deactivateMutation, { loading: loadingM2 }] = useMutation<
    DeactivateVersionResponse,
    VersionActionVars
  >(DEACTIVATE_VERSION, { onCompleted });
  const [deployMutation, { loading: loadingM3 }] = useMutation<
    DeployVersionResponse,
    VersionActionVars
  >(DEPLOY_VERSION, { onCompleted });
  const [stopMutation, { loading: loadingM4 }] = useMutation<
    StopVersionResponse,
    VersionActionVars
  >(STOP_VERSION, { onCompleted });

  const mutationLoading = [loadingM1, loadingM2, loadingM3, loadingM4].some(
    el => el
  );

  function getMutationVars(versionId: string, comment?: string) {
    const variables = {
      variables: {
        input: {
          versionId: versionId
        }
      }
    } as any;

    if (comment) {
      variables.variables.input.comment = comment;
    }

    return variables;
  }

  return {
    [actions.activate]: activateMutation,
    [actions.deploy]: deployMutation,
    [actions.stop]: stopMutation,
    [actions.deactivate]: deactivateMutation,
    getMutationVars,
    mutationLoading
  };
}
