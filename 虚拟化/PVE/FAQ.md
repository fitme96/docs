1.  TASK ERROR: activating LV ‘pve/data’ failed: Activation of logical volume pve/data is prohibited while logical volume pve/data_tmeta is active.

lvchange -an pve/data_tmeta
lvchange -an pve/data_tdata
lvchange -ay pve

参考：https://blog.csdn.net/feitianyul/article/details/125417765