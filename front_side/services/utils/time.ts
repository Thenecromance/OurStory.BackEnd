

class TimeFormatter {
    date: Date;

    constructor(unix_: number | string | Date) {
        if (typeof unix_ === 'number') {

            this.date = new Date(unix_);
        } else if (typeof unix_ == 'string') {
            this.date = new Date(unix_);

        } else if (unix_ instanceof Date) {
            this.date = unix_;
        }
    }

    // Convert date to unix timestamp in seconds
    toUnix(): number {
        return this.date.getTime() / 1000;
    }

    // Convert date to Date object
    toDate(): Date {
        return this.date;
    }

    // Convert date to string in format 'YYYY-MM-DD'
    toDateString(): string {
        return this.date.toISOString().split('T')[0];
    }

    // Convert date to string in format 'YYYY-MM-DD HH:MM:SS'
    toDateTimeString(): string {
        return this.date.toISOString().replace('T', ' ').split('.')[0];
    }
}


