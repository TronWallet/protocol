package org.tron.core.db;

import com.google.protobuf.ByteString;
import org.joda.time.DateTime;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.tron.common.utils.ByteArray;

import java.util.Optional;

public class DynamicPropertiesStore extends TronDatabase {

  private static final Logger logger = LoggerFactory.getLogger("DynamicPropertiesStore");

  private static final long MAINTENANCE_TIME_INTERVAL = 24 * 3600 * 1000;// (ms)

  private static final byte[] LATEST_BLOCK_HEADER_TIMESTAMP = "latest_block_header_timestamp"
      .getBytes();
  private static final byte[] LATEST_BLOCK_HEADER_NUMBER = "latest_block_header_number".getBytes();
  private static final byte[] LATEST_BLOCK_HEADER_HASH = "latest_block_header_hash".getBytes();
  private static final byte[] STATE_FLAG = "state_flag"
      .getBytes();// 1 : is maintenance, 0 : is not maintenance


  private BlockFilledSlots blockFilledSlots = new BlockFilledSlots();
  private DateTime nextMaintenanceTime;

  private DynamicPropertiesStore(String dbName) {
    super(dbName);

    try {
      this.getLatestBlockHeaderTimestamp();
    } catch (IllegalArgumentException e) {
      this.saveLatestBlockHeaderTimestamp(0);
    }

    try {
      this.getLatestBlockHeaderNumber();
    } catch (IllegalArgumentException e) {
      this.saveLatestBlockHeaderNumber(0);
    }

    try {
      this.getLatestBlockHeaderHash();
    } catch (IllegalArgumentException e) {
      this.saveLatestBlockHeaderHash(ByteString.copyFrom(ByteArray.fromHexString("00")));
    }

    try {
      this.getStateFlag();
    } catch (IllegalArgumentException e) {
      this.saveStateFlag(0);
    }

  }

  @Override
  public void put(byte[] key, Object item) {

  }

  @Override
  public void delete(byte[] key) {

  }

  @Override
  public Object get(byte[] key) {
    return null;
  }

  @Override
  public boolean has(byte[] key) {
    return false;
  }

  private static DynamicPropertiesStore instance;

  /**
   * create fun.
   *
   * @param dbName the name of database
   */
  public static DynamicPropertiesStore create(String dbName) {
    if (instance == null) {
      synchronized (DynamicPropertiesStore.class) {
        if (instance == null) {
          instance = new DynamicPropertiesStore(dbName);
        }
      }
    }
    return instance;
  }


  /**
   * get timestamp of creating global latest block.
   */
  public long getLatestBlockHeaderTimestamp() {
    byte[] t = this.dbSource.getData(LATEST_BLOCK_HEADER_TIMESTAMP);
    return Optional.ofNullable(t)
            .map(ByteArray::toLong)
            .orElseThrow(() -> new IllegalArgumentException("not found latest block header timestamp"));
  }

  /**
   * get number of global latest block.
   */
  public long getLatestBlockHeaderNumber() {
    byte[] n = this.dbSource.getData(LATEST_BLOCK_HEADER_NUMBER);
    return Optional.ofNullable(n)
            .map(ByteArray::toLong)
            .orElseThrow(() -> new IllegalArgumentException("not found latest block header number"));
  }

  public int getStateFlag() {
    byte[] n = this.dbSource.getData(STATE_FLAG);

    if (n == null || n.length == 0) {
      throw new IllegalArgumentException("not found maintenance flag");
    }

    return ByteArray.toInt(n);
  }

  /**
   * get id of global latest block.
   */
  public ByteString getLatestBlockHeaderHash() {
    byte[] h = this.dbSource.getData(LATEST_BLOCK_HEADER_HASH);
    return Optional.ofNullable(h)
            .map(ByteString::copyFrom)
            .orElseThrow(() -> new IllegalArgumentException("not found latest block header id"));
  }

  /**
   * save timestamp of creating global latest block.
   */
  public void saveLatestBlockHeaderTimestamp(long t) {
    logger.info("update latest block header timestamp = {}", t);
    this.dbSource.putData(LATEST_BLOCK_HEADER_TIMESTAMP, ByteArray.fromLong(t));
  }

  /**
   * save number of global latest block.
   */
  public void saveLatestBlockHeaderNumber(long n) {
    logger.info("update latest block header number = {}", n);
    this.dbSource.putData(LATEST_BLOCK_HEADER_NUMBER, ByteArray.fromLong(n));
  }

  /**
   * save id of global latest block.
   */
  public void saveLatestBlockHeaderHash(ByteString h) {
    logger.info("update latest block header id = {}", ByteArray.toHexString(h.toByteArray()));
    this.dbSource.putData(LATEST_BLOCK_HEADER_HASH, h.toByteArray());
  }

  public void saveStateFlag(int n) {
    logger.info("update state flag = {}", n);
    this.dbSource.putData(STATE_FLAG, ByteArray.fromInt(n));
  }

  public BlockFilledSlots getBlockFilledSlots() {
    return blockFilledSlots;
  }


  public DateTime getNextMaintenanceTime() {
    return nextMaintenanceTime;
  }

  public void updateMaintenanceTime(){
    nextMaintenanceTime = nextMaintenanceTime.plus(MAINTENANCE_TIME_INTERVAL);
  }

}