access(all) contract PseudoRandomGenerator {

    access(all) let StoragePath: StoragePath
    access(all) let PublicPath: PublicPath

    access(all) event NextUInt64(id: UInt64, value: UInt64)

    /// While not limited to 128 bits of state, this PRG is largely informed by XORShift128+
    ///
    access(all) resource PRG {
        access(all) let sourceOfRandomness: [UInt8]
        access(all) let salt: Word64
        
        /// The states below are of type Word64 to prevent overflow/underflow
        //
        access(all) var state0: Word64
        access(all) var state1: Word64
        
        init(sourceOfRandomness: [UInt8], salt: UInt64) {
            pre {
                sourceOfRandomness.length == 32: "Expecting 32 bytes"
            }
            self.sourceOfRandomness = sourceOfRandomness
            self.salt = Word64(salt)

            // Convert the seed bytes to two Word64 values for state initialization
            let segment0 = PseudoRandomGenerator.bytesToWord64(bytes: sourceOfRandomness, start: 0)
            let segment1 = PseudoRandomGenerator.bytesToWord64(bytes: sourceOfRandomness, start: 8)
            let segment2 = PseudoRandomGenerator.bytesToWord64(bytes: sourceOfRandomness, start: 16)
            let segment3 = PseudoRandomGenerator.bytesToWord64(bytes: sourceOfRandomness, start: 24)

            self.state0 = segment0 ^ segment1
            self.state1 = segment2 ^ segment3
        }

        access(all) fun nextUInt64(): UInt64 {
            var a: Word64 = self.state0
            let b: Word64 = self.state1

            self.state0 = b
            a = a ^ (a << 23) // a
            a = a ^ (a >> 17) // b
            a = b ^ (b >> 5) // c
            self.state1 = a

            let randUInt64: UInt64 = UInt64((a + b) ^ self.salt)
            emit NextUInt64(id: self.uuid, value: randUInt64)
            return randUInt64
        }
    }

    /// Helper function to convert an array of bytes to Word64
    access(contract) fun bytesToWord64(bytes: [UInt8], start: Int): Word64 {
        pre {
            start % 8 == 0: "Expecting start to be a multiple of 8"
            bytes.length % 8 == 0: "byte array length to be a multiple of 8"
        }
        var value: UInt64 = 0
        var i = 0
        while i < 8 {
            value = value << 8 | UInt64(bytes[start + i])
            i = i + 1
        }
        return Word64(value)
    }

    access(all) fun createPRG(sourceOfRandomness: [UInt8], salt: UInt64): @PRG {
        return <- create PRG(sourceOfRandomness: sourceOfRandomness, salt: salt)
    }

    init() {
        self.StoragePath = StoragePath(identifier: "PseudoRandomGenerator_".concat(self.account.address.toString()))!
        self.PublicPath = PublicPath(identifier: "PseudoRandomGenerator_".concat(self.account.address.toString()))!
    }
}
