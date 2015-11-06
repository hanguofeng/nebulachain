package main

import (
	"os"

	"log"
)

type SyncManager struct {
	plans []*Plan
}

func NewSyncManager() *SyncManager {
	mgr := &SyncManager{}
	mgr.plans = []*Plan{}
	return mgr
}

func (this *SyncManager) AddPlan(plan *Plan) {
	this.plans = append(this.plans, plan)
}

func (this *SyncManager) Work() {
	for _, plan := range this.plans {
		this.workOnPlan(plan)
	}
}

func (this *SyncManager) workOnPlan(plan *Plan) {
	util := &StorageUtil{Bucket: plan.Bucket, Region: plan.Region}
	err := util.init()

	if err != nil {
		log.Fatalf("connect to stroage error : %s", err.Error())
		os.Exit(2)
	}

	walker := &FileWalker{
		entrypoint: plan.Path,
	}

	walker.Walk(func(fullname string, shortname string, fp string) {
		shortname = plan.KeyPrefix + "/" + shortname
		log.Printf("[INFO] syncing %s(%s)", shortname, fullname)
		info, err := util.getFileInfo(shortname)
		remoteFp := info.Metadata["Fingerprint"]
		if err != nil || *remoteFp != fp {
			result, err := util.upload(shortname, fullname, fp)
			log.Printf("[INFO] synced %s,%s", result, err)
		} else {
			log.Printf("[INFO] omited %s", fullname)

		}
	})
}
