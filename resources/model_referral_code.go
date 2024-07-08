/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ReferralCode struct {
	// Referral code itself, unique identifier
	Id string `json:"id"`
	// Status of the code, belonging to this user (referrer):   1. infinity: the code have unlimited usage count and user can get points for each user who scanned passport   2. active: the code is not used yet by another user (referee)   3. awaiting: the code is used by referee who has scanned passport, but the referrer hasn't yet   4. rewarded: the code is used, both referee and referrer have scanned passports   5. consumed: the code is used by referee who has not scanned passport yet  The list is sorted by priority. E.g. if the referee has scanned passport, but referrer not, the status would be `consumed`. If both not scann passport yet status would be `awaiting`.
	Status string `json:"status"`
}
