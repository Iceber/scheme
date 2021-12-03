# Global Scheme Demo

以插件的形式注册任何资源的 scheme 到 global scheme 中，主要用于 CRD 资源的结构注册

```sh
$ make run
go build --buildmode=plugin -o ./plugins/karmada.so ./plugins/karmada
go build --buildmode=plugin -o ./plugins/clusterapi.so ./plugins/clusterapi
go build -o globalscheme main.go
./globalscheme
register karmada
register karmada
+-------------------------------+----------------------------+-------------------------------------+
|          APIVersion           |            Kind            |               Struct                |
+-------------------------------+----------------------------+-------------------------------------+
|  cluster.karmada.io/v1alpha1  |       CreateOptions        |          v1.CreateOptions           |
|              v1               |        APIVersions         |           v1.APIVersions            |
|              v1               |          APIGroup          |             v1.APIGroup             |
|   cluster.x-k8s.io/v1beta1    |   MachineDeploymentList    |    v1beta1.MachineDeploymentList    |
|   work.karmada.io/v1alpha1    |         GetOptions         |            v1.GetOptions            |
|  cluster.karmada.io/v1alpha1  |         WatchEvent         |            v1.WatchEvent            |
|  cluster.karmada.io/v1alpha1  |       DeleteOptions        |          v1.DeleteOptions           |
|   cluster.x-k8s.io/v1beta1    |       DeleteOptions        |          v1.DeleteOptions           |
|   cluster.x-k8s.io/v1beta1    |       CreateOptions        |          v1.CreateOptions           |
|   cluster.x-k8s.io/v1beta1    |        PatchOptions        |           v1.PatchOptions           |
|   work.karmada.io/v1alpha1    | ClusterResourceBindingList | v1alpha1.ClusterResourceBindingList |
|  work.karmada.io/__internal   |         WatchEvent         |          v1.InternalEvent           |
|  cluster.karmada.io/v1alpha1  |        PatchOptions        |           v1.PatchOptions           |
|   cluster.x-k8s.io/v1beta1    |          Cluster           |           v1beta1.Cluster           |
|   cluster.x-k8s.io/v1beta1    |       UpdateOptions        |          v1.UpdateOptions           |
|   work.karmada.io/v1alpha1    |   ClusterResourceBinding   |   v1alpha1.ClusterResourceBinding   |
|   work.karmada.io/v1alpha1    |      ResourceBinding       |      v1alpha1.ResourceBinding       |
|   cluster.x-k8s.io/v1beta1    |        ListOptions         |           v1.ListOptions            |
|              v1               |           Status           |              v1.Status              |
|   work.karmada.io/v1alpha1    |         WatchEvent         |            v1.WatchEvent            |
|   work.karmada.io/v1alpha1    |       CreateOptions        |          v1.CreateOptions           |
|   work.karmada.io/v1alpha1    |        PatchOptions        |           v1.PatchOptions           |
|   cluster.x-k8s.io/v1beta1    |         MachineSet         |         v1beta1.MachineSet          |
|   work.karmada.io/v1alpha1    |    ResourceBindingList     |    v1alpha1.ResourceBindingList     |
| cluster.karmada.io/__internal |         WatchEvent         |          v1.InternalEvent           |
|  cluster.karmada.io/v1alpha1  |       UpdateOptions        |          v1.UpdateOptions           |
|   cluster.x-k8s.io/v1beta1    |        MachineList         |         v1beta1.MachineList         |
|   cluster.x-k8s.io/v1beta1    |     MachineHealthCheck     |     v1beta1.MachineHealthCheck      |
|   cluster.x-k8s.io/v1beta1    |       MachineSetList       |       v1beta1.MachineSetList        |
|   work.karmada.io/v1alpha1    |       DeleteOptions        |          v1.DeleteOptions           |
|   work.karmada.io/v1alpha1    |       UpdateOptions        |          v1.UpdateOptions           |
|  cluster.karmada.io/v1alpha1  |        ClusterList         |        v1alpha1.ClusterList         |
|              v1               |        APIGroupList        |           v1.APIGroupList           |
|   cluster.x-k8s.io/v1beta1    |        ClusterClass        |        v1beta1.ClusterClass         |
|   work.karmada.io/v1alpha1    |          WorkList          |          v1alpha1.WorkList          |
|  cluster.karmada.io/v1alpha1  |        ListOptions         |           v1.ListOptions            |
|   cluster.x-k8s.io/v1beta1    |          Machine           |           v1beta1.Machine           |
|   cluster.x-k8s.io/v1beta1    |     MachineDeployment      |      v1beta1.MachineDeployment      |
|   cluster.x-k8s.io/v1beta1    |        ClusterList         |         v1beta1.ClusterList         |
|   cluster.x-k8s.io/v1beta1    |         WatchEvent         |            v1.WatchEvent            |
|  cluster.x-k8s.io/__internal  |         WatchEvent         |          v1.InternalEvent           |
|   cluster.x-k8s.io/v1beta1    |         GetOptions         |            v1.GetOptions            |
|              v1               |      APIResourceList       |         v1.APIResourceList          |
|   cluster.x-k8s.io/v1beta1    |      ClusterClassList      |      v1beta1.ClusterClassList       |
|   cluster.x-k8s.io/v1beta1    |   MachineHealthCheckList   |   v1beta1.MachineHealthCheckList    |
|   work.karmada.io/v1alpha1    |            Work            |            v1alpha1.Work            |
|   work.karmada.io/v1alpha1    |        ListOptions         |           v1.ListOptions            |
|  cluster.karmada.io/v1alpha1  |          Cluster           |          v1alpha1.Cluster           |
|  cluster.karmada.io/v1alpha1  |         GetOptions         |            v1.GetOptions            |
+-------------------------------+----------------------------+-------------------------------------+

```
